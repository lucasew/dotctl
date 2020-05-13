package main

import (
	"github.com/spf13/cobra"
)

func init() {
	bringCmd := &cobra.Command{
		Use:   "bring",
		Short: "bring everything described in the rules file to the repository",
		Long:  "bring everything described in the rules file to the repository",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := DotCtlRepository.Bring(args...)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
		},
	}
	RootCMD.AddCommand(bringCmd)
}
