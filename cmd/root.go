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
	// Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// var input string
		if showMsgs {

			ans := ShowThreeMsgs()
			// for i, msg := range ans {
			// 	fmt.Printf("%s pick %d\n", msg, i)

			// }
			// // for i := range 3 {
			// // 	fmt.Printf("%s pick %d\n", ans[i+1], i+1)
			// // }
			// fmt.Printf("pick the number\n")
			// var in int
			// fmt.Scan(&in)
			// str := ans[in]

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
