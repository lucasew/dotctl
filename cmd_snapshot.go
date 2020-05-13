package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	snapshotCmd := &cobra.Command{
		Use:   "snapshot",
		Short: "import the files based on the rules and commit automatically",
		Long:  "import the files based on the rules and commit automatically",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := DotCtlRepository.Git("fetch", "--all")
			if err != nil {
				cmd.PrintErrln(err)
			}
			hostname := getHostname()
			err = DotCtlRepository.Git("switch", "-c", hostname)
			if err != nil {
				cmd.PrintErrln(err)
			}
			err = DotCtlRepository.Bring(args...)
			if err != nil {
				cmd.PrintErrln(err)
			}
			err = DotCtlRepository.Git("add", "-A")
			if err != nil {
				cmd.PrintErrln(err)
			}
			err = DotCtlRepository.Git("commit", "-m", time.Now().String())
		},
	}
	RootCMD.AddCommand(snapshotCmd)
}

func getHostname() string {
	hostname, ok := os.LookupEnv("HOSTNAME")
	if ok {
		return hostname
	}
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}
