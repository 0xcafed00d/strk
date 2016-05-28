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

	repoChan := make(chan string)
	commitsChan := make(chan repoCommits)

	for n := 0; n < 8; n++ {
		go func(n int) {
			log.Println("Staring go func", n)
			for name := <-repoChan; name != ""; name = <-repoChan {
				log.Println("Fetching commits for repo: ", name)
				commits, err := streak.GetCommits("simulatedsimian", name, os.Args[1])
				commitsChan <- repoCommits{name, commits, err}
			}
			log.Println("End go func", n)
		}(n)
	}

	repos, err := streak.GetRepos("simulatedsimian", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	(*repos)[4].Name = "wibble"

	go func() {
		for _, v := range *repos {
			repoChan <- v.Name
		}
		close(repoChan)
	}()

	for n := 0; n < len(*repos); n++ {
		res := <-commitsChan
		if res.err != nil {
			log.Println(res.repoName, res.err)
		} else {
			log.Println(res.repoName, len(*res.CommitsJSON))
		}
	}
}
