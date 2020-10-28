package inspectjwt

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"reflect"
	"strings"
)

var colors = []string{"\033[35m", "\033[33m"}

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type appEnv struct {
	jwt string
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("inspect-jwt", flag.ContinueOnError)
	tokenString := fl.String("t", "", "The token to inspect")

	if err := fl.Parse(args); err != nil {
		return err
	}
	app.jwt = strings.TrimSpace(*tokenString)

	if app.jwt == "" {
		_, _ = fmt.Fprintln(os.Stderr, "got an empty jwt!")
		fl.Usage()
		return flag.ErrHelp
	}
	return nil
}

func (app *appEnv) run() error {
	simpleToken, err := parseToken(app.jwt)
	if err != nil {
		return err
	}

	reflectedFields := reflect.ValueOf(*simpleToken)

	for i := 0; i < reflectedFields.NumField(); i++ {
		prettyPrintJson(colors[i], reflectedFields.Field(i).Interface())
	}
	return nil
}

func parseToken(tokenString string) (*simpleToken, error) {
	parser := jwt.Parser{SkipClaimsValidation: true, UseJSONNumber: true}
	token, _, err := parser.ParseUnverified(tokenString, &jwt.MapClaims{})
	if err != nil {
		return nil, err
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
