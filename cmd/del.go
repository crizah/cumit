package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var delcmd = &cobra.Command{
	Use: "del",
	Aliases: []string{
		"delete",
	},

	Short: "delete a picture",
	Long:  "delete /home/crizah/Downloads/jpeg",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("deleting %s, %s", args[0], DeletePix(args[0]))

	},
}

func init() {
	rootcmd.AddCommand(delcmd)
}
