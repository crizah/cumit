package cmd

import "math/rand"

var showMsgs bool
var msgs = []string{"poopy butthead", "daddy dookie", "she moves with a purpose",
	"what a magnificent purpose", "is this hell be fr rn", "ill take all ur thanks and all ur sympathies",
	"cry for me", "plenty of time here to nurture concequence", "its a bit of a joke"}

const (
	MAX = 4
	MIN = 0
)

func ShowThreeMsgs() map[int]string {
	// should be able to select one of the three messages
	ans := make(map[int]string)
	for i := range 3 {
		x := rand.Intn(MAX-MIN) + MIN // from min-max
		ans[i+1] = msgs[x]

	}
	return ans

}

func init() {
	rootcmd.Flags().BoolVarP(&showMsgs, "show", "s", false, "shows 3 random messages")
}
