package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// git commit -> prepare-commit-msg hook -> cumit runs -> user picks message -> git uses it
var rootcmd = &cobra.Command{
	Use:   "cumit",
	Short: "Cumit is a cli tool to generate silly git commit msgs",
	Long:  "thats it, that what it is",

	Run: func(cmd *cobra.Command, args []string) {

		if showMsgs {

			ans := ShowThreeMsgs()
			ans = append(ans, "refresh")

			var str string
			prompt := &survey.Select{
				Message: "Pick one",
				Options: ans,
			}
			survey.AskOne(prompt, &str)
			if str == "refresh" {
				ans = ShowThreeMsgs()
				prompt := &survey.Select{
					Message: "Pick one",
					Options: ans,
				}
				survey.AskOne(prompt, &str)
			}
			fmt.Printf("chosen: %s\n", str)

			return
		}

	},
}

func init() {
	rootcmd.AddCommand(initcmd)
	rootcmd.AddCommand(hookcmd)
	rootcmd.Flags().BoolVarP(&showMsgs, "show", "s", false, "shows 3 random messages")
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "u piece of shit %s", err)
		os.Exit(1)
	}

}
