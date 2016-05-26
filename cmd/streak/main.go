package main

import (
	"fmt"
	"os"

	"github.com/simulatedsimian/strk/streak"
)

func main() {

	repos, err := streak.GetRepos("simulatedsimian", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range *repos {
		commits, err := streak.GetCommits("simulatedsimian", v.Name, os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(v.Name, len(*commits))
	}
}
