package main

import (
	"os"
	"runtime"
)

func getEditor() string {
    ret, ok := os.LookupEnv("EDITOR")
    if ok {
        return ret
    }
    switch runtime.GOOS {
    case "windows":
        return "notepad"
    default:
        return "nano"
    }
}
