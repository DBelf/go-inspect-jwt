package main

import (
	"github.com/dbelf/go-jwt-inspect/inspectjwt"
	"os"
)

func main() {
	os.Exit(inspectjwt.CLI(os.Args[1:]))
}
