package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type DotctlRepo struct {
    basepath string
}

// NewDotctlRepo this function is not side effect free: it uses environment variables
func NewDotctlRepo() DotctlRepo {
    dir, ok := os.LookupEnv("DOTCTL_DIR")
    if !ok {
        dir = "~/.dotctl"
    }
    dir, err := homedir.Expand(dir)
    if err != nil {
        panic(err)
    }
    err = os.MkdirAll(dir, os.ModePerm)
    if err != nil {
        panic(err) // I want it to cancel everything if it cant create dir
    }
    cmd := exec.Command("git", "init")
    cmd.Dir = dir
    err = cmd.Run()
    if err != nil {
        panic(err)
    }
    return DotctlRepo{dir}
}

func (d DotctlRepo) DotCtlRepo() string {
    return d.basepath
}

func (d DotctlRepo) JoinPath(path string) string {
    return filepath.Join(d.basepath, path)
}

// RulesRepository builder for RuleRepository
func (d DotctlRepo) RulesRepository() RuleRepository {
    return NewRuleRepository(d.JoinPath("files.txt"))
}
