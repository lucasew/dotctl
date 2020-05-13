package main

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
    rulesCmd := &cobra.Command{
        Use: "rules",
        Short: "administration of dotctl file rules for files to add",
        Long: "administation for rules for files that dotctl will use",
        Run: func(cmd *cobra.Command, args []string) {
            cmd.Help()
        },
    }
    rulesCmd.AddCommand(&cobra.Command{
        Use: "ls",
        Short: "show all rules",
        Long: "show all rules",
        Run: func (cmd *cobra.Command, args []string) {
            rules := DotCtlRepository.RulesRepository().GetRules()
            for _, rule := range rules {
                println(rule)
            }
        },
    })
    rulesCmd.AddCommand(&cobra.Command{
        Use: "add",
        Short: "add one or more rules",
        Long: "add one or more rules to dotctl",
        Args: cobra.ArbitraryArgs,
        Run: func(cmd *cobra.Command, args []string) {
            rulesRepo := DotCtlRepository.RulesRepository()
            err := rulesRepo.AppendRules(args...)
            if err != nil {
                cmd.PrintErrln(err)
                return
            }
        },
    })
    rulesCmd.AddCommand(&cobra.Command{
        Use: "edit",
        Short: "call $EDITOR to edit the rules file",
        Long: "call $EDITOR to edit the rules file",
        Run: func(cmd *cobra.Command, args []string) {
            filename := DotCtlRepository.RulesRepository().RulesFilename()
            editor := getEditor()
            ecmd := exec.Command(editor, filename)
            ecmd.Stdin = os.Stdin
            ecmd.Stdout = os.Stdout
            ecmd.Stderr = os.Stderr
            err := ecmd.Run()
            if err != nil {
                cmd.PrintErrln(err)
                return
            }
        },
    })
    RootCMD.AddCommand(rulesCmd)
}
