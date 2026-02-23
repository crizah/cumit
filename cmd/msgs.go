package cmd

import (
	"cumit/msg"
	"math/rand"
)

var showMsgs bool

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

		ans = append(ans, msg.MSGS[x])

	}
	return ans

}
