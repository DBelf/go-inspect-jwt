package inspectjwt

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"reflect"
	"strings"
)

var colors = []string{"\033[35m", "\033[33m"}

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 1
	}
	if err = app.run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error while running app: %v\n", err)
		return 1
	}
	return 0
}

type appEnv struct {
	jwt          string
	checkExpired bool
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("inspect-jwt", flag.ContinueOnError)
	tokenString := fl.String("t", "", "The token to inspect")
	isExpiredCheck := fl.Bool("exp", false, "Check whether the token is expired")

	if err := fl.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}
	app.jwt = strings.TrimSpace(*tokenString)
	app.checkExpired = *isExpiredCheck

	if app.jwt == "" {
		fl.Usage()
	}
	return nil
}

func (app *appEnv) run() error {
	simpleToken, err := parseToken(app.jwt)
	if err != nil {
		return fmt.Errorf("failed to read token: %w", err)
	}

	reflectedFields := reflect.ValueOf(*simpleToken)
	if app.checkExpired {
		printTokenIsExpired(*simpleToken)
	}

	for i := 0; i < reflectedFields.NumField(); i++ {
		prettyPrintJson(colors[i], reflectedFields.Field(i).Interface())
	}
	return nil
}

func printTokenIsExpired(token simpleToken) {
	valid := token.Claims.Valid()
	if valid == nil {
		fmt.Println("Token is still valid.")
	} else {
		fmt.Println("Token is invalid.")
	}
}

func parseToken(tokenString string) (*simpleToken, error) {
	parser := jwt.Parser{SkipClaimsValidation: true, UseJSONNumber: true}
	token, _, err := parser.ParseUnverified(tokenString, &jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}
	return tokenToSimpleToken(token), nil
}

func prettyPrintJson(color string, data interface{}) {
	prettyOutput, _ := json.MarshalIndent(data, "", "  ")
	fmt.Print(color, string(prettyOutput))
}

type simpleToken struct {
	Header map[string]interface{}
	Claims jwt.Claims
}

func tokenToSimpleToken(token *jwt.Token) *simpleToken {
	return &simpleToken{Claims: token.Claims, Header: token.Header}
}
