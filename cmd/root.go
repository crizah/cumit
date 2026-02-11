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
	Run: func(cmd *cobra.Command, args []string) {
		if showMsgs {
			fmt.Println("pick one")
			ans := ShowThreeMsgs()
			for i := range 3 {
				fmt.Printf("%s\n", ans[i])
			}
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
