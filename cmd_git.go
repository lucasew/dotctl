package main

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func (r DotctlRepo) Git(args ...string) error {
	ecmd := exec.Command("git", args...)
	ecmd.Stdin = os.Stdin
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	ecmd.Dir = r.DotCtlRepo()
	return ecmd.Run()
}

func init() {
	gitCmd := &cobra.Command{
		Use:   "git",
		Short: "use git in the dotctl repo from any location",
		Long:  "use git in the dotctl from any location",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := DotCtlRepository.Git(args...)
			if err != nil {
				cmd.PrintErrln(err.Error())
			}
		},
	}
	RootCMD.AddCommand(gitCmd)
}
