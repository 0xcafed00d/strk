package main

import (
	"fmt"
	"os"

	"github.com/simulatedsimian/strk/streak"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: streak <username> <apikey>")
		os.Exit(1)
	}
	key := os.Args[2]
	username := os.Args[1]

	streaks, err := streak.GetStreaks(username, key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	longest := streak.LongestStreak(streaks)
	fmt.Println("Longest Streak: ", longest)

	user := streak.UserPatchBioJSON{Bio: "Longest Streak: " + longest.String()}
	err = streak.UpdateUser(user, key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
