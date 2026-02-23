package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

		path := ".git/hooks/prepare-commit-msg"
		ex, err := getPath()
		if err != nil {
			fmt.Fprintf(os.Stderr, "cabt get executable path %s\n", err)
			os.Exit(1)
		}

		content := fmt.Sprintf(`#!/bin/sh
# cumit hook
GIT_EDITOR=true "%s" hook "$1" "$2"
`, ex)

		err = os.WriteFile(path, []byte(content), 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "coulrnt write hook %s\n", err)
			os.Exit(1)
		}

		c := exec.Command("git", "config", "core.editor", "true")
		if err := c.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to set editor: %s\n", err)
		}

		fmt.Println("hell yeah")

	},
}

func getPath() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err

	}

	return path, nil
}

func checkGitRepo() bool {
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		return false

	}
	return true

}
