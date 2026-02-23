package cmd

import (
	"math/rand"

	"github.com/crizah/cumit/msg"
)

var showMsgs bool

const (
	MIN = 0
)

func ShowThreeMsgs() []string {
	// should be able to select one of the three messages
	MAX := len(msg.MSGS)
	ans := make([]string, 0)
	put := make(map[int]bool)
	for range 3 {
		// dont pick the one previosuly picked
		var x int
		a := rand.Intn(MAX-MIN) + MIN
		if put[a] {
			// choose another
			a = rand.Intn(MAX-MIN) + MIN
			for put[a] {
				a = rand.Intn(MAX-MIN) + MIN

			}

		}
		x = a
		put[x] = true

		ans = append(ans, msg.MSGS[x])

	}
	return ans

}
