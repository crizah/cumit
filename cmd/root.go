package cmd

import (
	"fmt"
	"os"

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
			fmt.Println("pick one")
			ans := ShowThreeMsgs()
			for i, msg := range ans {
				fmt.Printf("%s pick %d\n", msg, i)

			}
			// for i := range 3 {
			// 	fmt.Printf("%s pick %d\n", ans[i+1], i+1)
			// }
			fmt.Printf("pick the number\n")
			var in int
			fmt.Scan(&in)
			str := ans[in]
			fmt.Printf("chosen %s\n", str)

			return
		}

	},
}

// Source - https://stackoverflow.com/a/41761404
// Posted by Shahriar, modified by community. See post 'Timeline' for change history
// Retrieved 2026-02-20, License - CC BY-SA 3.0

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "u piece of shit %s", err)
		os.Exit(1)
	}

}
