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
}

func main() {

	repoChan := make(chan string)
	commitsChan := make(chan repoCommits)

	go func() {
		repos, err := streak.GetRepos("simulatedsimian", os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, v := range *repos {
			repoChan <- v.Name
		}
		close(repoChan)
	}()

	for n := 0; n < 16; n++ {
		go func(n int) {
			log.Println("Staring go func", n)
			for {
				name := <-repoChan
				if name == "" {
					break
				}
				log.Println("Fetching commits for repo: ", name)

				commits, err := streak.GetCommits("simulatedsimian", name, os.Args[1])
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				commitsChan <- repoCommits{name, commits}
			}
			log.Println("End go func", n)
		}(n)
	}

	for {
		res := <-commitsChan
		log.Println(res.repoName, len(*res.CommitsJSON))
	}
}
