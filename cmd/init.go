package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initcmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise",
	Long:  "install hook into .git/hooks/prepare-commit-msg",
	Run: func(cmd *cobra.Command, args []string) {
		// check if .git
		check := checkGitRepo()
		if !check {
			fmt.Printf("Not a git repo dumdum\n")
			os.Exit(1)
		}

		path := "../.git/hooks/prepare-commit-msg"

		content := `
		#!/bin/sh
# cumit hook
cumit hook "$1" "$2"`

		err := os.WriteFile(path, []byte(content), 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "coulrnt write hook %s\n", err)
			os.Exit(1)
		}

		fmt.Println("hell yeah")

	},
}

func checkGitRepo() bool {
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		return false

	}
	return true

}
