package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

var DotCtlRepository = NewDotctlRepo()

func main() {
	err := os.Chdir(DotCtlRepository.basepath)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}
	godotenv.Load() // if .env dont exists, no problem
	err = RootCMD.Execute()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}
