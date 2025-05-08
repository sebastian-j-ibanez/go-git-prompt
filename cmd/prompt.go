/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sebastian-j-ibanez/go-git-prompt/status"
	"github.com/spf13/cobra"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Output a Git status prompt for the current directory.",
	Long:  `Output a Git status prompt for the current directory. This command should be placed in your shell prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt, err := status.GetGitInfo()
		if !prompt.IsEmpty() && err == nil {
			prompt.Print()
		}
	},
}

func init() {
	rootCmd.AddCommand(promptCmd) // Add a nerd icons flag
}
