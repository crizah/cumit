package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{
	Use:   "cumit",
	Short: "Cumit is a cli tool to generate silly git commit msgs",
	Long:  "thats it, that what it is",

	Run: func(cmd *cobra.Command, args []string) {

		if showMsgs {

			ans := ShowThreeMsgs()

			var str string
			prompt := &survey.Select{
				Message: "Pick one",
				Options: ans,
			}
			survey.AskOne(prompt, &str)

			fmt.Printf("chosen: %s\n", str)

			return
		}

	},
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "u piece of shit %s", err)
		os.Exit(1)
	}

}
