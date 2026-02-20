package cmd

import (
	"log"
	"os/exec"
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
