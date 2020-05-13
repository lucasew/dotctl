package main

import (
	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
    Use: "dotctl",
    Short: "dotctl helps you to organize your dotfiles",
    Long: "dotctl helps you to organize your dotfiles copying them to a folder for easy git administration",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}
