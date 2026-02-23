package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var hookcmd = &cobra.Command{
	Use:   "hook [commit-msg-file] [commit-source]",
	Short: "make the hook",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commitFile := args[0]

		// args[1] is -m
		// if -m is used, return
		if len(args) > 1 && args[1] != "" {
			return
		}

		msgs := ShowThreeMsgs()
		tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
		if err != nil {
			return
		}
		defer tty.Close()
		var str string
		prompt := &survey.Select{
			Message: "Pick one",
			Options: msgs,
		}
		err = survey.AskOne(prompt, &str)
		if err != nil {
			return

		}

		// write str into the commit file

		err = os.WriteFile(commitFile, []byte(str+"\n"), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to write commit msg %s\n", err)
			os.Exit(1)
		}
	},
}
