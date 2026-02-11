package cmd

import (
	"log"
	"math/rand"
	"os/exec"
)

var msgs = []string{"poopy butthead", "daddy dookie", "she moves with a purpose", "what a magnificent purpose", "is this hell be fr rn"}

const (
	MAX = 4
	MIN = 0
)

func CumitInit() {
	// init the commit
	// idk how to do that
	// how2do that
	// will only work in a repo with a .git in it
	// once initialised, when user does
	// "git commit" enter -> runs ShowAllMsgs()
	// user can pick any of the 3 (up and down keys)
	// on enter, user can add their own commit message also as one of the options
	// on enter, this function runs git commit <message>

}

func CommitMessage(message string) string {
	cmd := exec.Command("git", "commit", "-m", message)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		return "nah"
	}

	return "yay"

}

func DeletePix(path string) string {
	cmd := exec.Command("rm", path)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		return "nah"
	}

	return "yay"

}

func ShowThreeMsgs() []string {
	// should be able to select one of the three messages
	ans := make([]string, 0)
	for range 3 {
		x := rand.Intn(MAX-MIN) + MIN // from min-max
		ans = append(ans, msgs[x])

	}
	return ans

}
