package cmd

var showMsgs bool

// var msgscmd = &cobra.Command{
// 	Use:   "show",
// 	Short: "shows mesgs",
// 	Long:  "picks 3 random msgs",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("pick one")
// 		ans := ShowAllMsgs()
// 		for i := range 3 {
// 			fmt.Printf("%s\n", ans[i])

// 		}

// 	},
// }

func init() {
	rootcmd.Flags().BoolVarP(&showMsgs, "show", "s", false, "shows 3 random messages")
}
