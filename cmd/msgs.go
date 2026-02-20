package cmd

import "math/rand"

var showMsgs bool
var msgs = []string{"poopy butthead", "daddy dookie", "she moves with a purpose",
	"what a magnificent purpose", "is this hell be fr rn", "ill take all ur thanks and all ur sympathies",
	"cry for me", "plenty of time here to nurture concequence", "its a bit of a joke"}

const (
	MAX = 8
	MIN = 0
)

func ShowThreeMsgs() []string {
	// should be able to select one of the three messages
	ans := make([]string, 0)
	put := make(map[int]bool)
	for range 3 {
		// dont pick the one previosuly picked
		var x int
		a := rand.Intn(MAX-MIN) + MIN // from min-max
		if put[a] {
			// choose another
			a = rand.Intn(MAX-MIN) + MIN
			for put[a] {
				a = rand.Intn(MAX-MIN) + MIN

			}

		} else {
			x = a
			put[x] = true

		}

		ans = append(ans, msgs[x])

	}
	return ans

}

func init() {
	rootcmd.Flags().BoolVarP(&showMsgs, "show", "s", false, "shows 3 random messages")
}
