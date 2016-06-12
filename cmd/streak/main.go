package main

import (
	"fmt"
	"log"
	"os"

	"github.com/simulatedsimian/strk/streak"
)

type repoCommits struct {
	repoName string
	*streak.CommitsJSON
	err error
}

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: streak <username> <apikey>")
		os.Exit(1)
	}
	key := os.Args[2]
	username := os.Args[1]

	repoChan := make(chan string)
	commitsChan := make(chan repoCommits)

	for n := 0; n < 8; n++ {
		go func(n int) {
			for name := <-repoChan; name != ""; name = <-repoChan {
				log.Println("Fetching commits for repo: ", name)
				commits, err := streak.GetCommits(username, name, key)
				commitsChan <- repoCommits{name, commits, err}
			}
		}(n)
	}

	repos, err := streak.GetRepos("simulatedsimian", key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		for _, v := range *repos {
			repoChan <- v.Name
		}
		close(repoChan)
	}()

	m := make(map[string]repoCommits)

	for n := 0; n < len(*repos); n++ {
		res := <-commitsChan
		m[res.repoName] = res
	}

	sr := streak.StreakRecorder{}

	for _, v := range m {
		for i := range *v.CommitsJSON {
			if (*v.CommitsJSON)[i].Committer.Login == username {
				t := (*v.CommitsJSON)[i].Commit.Author.Date
				sr.AddCommit(t)
			}
		}
	}

	longest := streak.LongestStreak(sr.GetStreaks())
	fmt.Println("Longest Streak: ", longest)

	user := streak.UserPatchBioJSON{Bio: "Longest Streak: " + longest.String()}
	streak.UpdateUser(user, key)
}
