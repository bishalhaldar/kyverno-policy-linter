package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kyverno-policy-linter",
	Short: "kyverno policy linter CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kyverno Policy Linter v0.0.1")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
