package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

var DotCtlRepository = NewDotctlRepo()

func main() {
	godotenv.Load() // if .env dont exists, no problem
	err := RootCMD.Execute()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}
