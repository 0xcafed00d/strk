package streak

import "log"

type repoCommits struct {
	repoName string
	*CommitsJSON
	err error
}

func GetStreaks(username, key string) ([]Streak, error) {

	repoChan := make(chan string)
	commitsChan := make(chan repoCommits)

	for n := 0; n < 8; n++ {
		go func(n int) {
			for name := <-repoChan; name != ""; name = <-repoChan {
				log.Println("Fetching commits for repo: ", name)
				commits, err := GetCommits(username, name, key)
				commitsChan <- repoCommits{name, commits, err}
			}
		}(n)
	}

	repos, err := GetRepos(username, key)
	if err != nil {
		return nil, err
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

	sr := StreakRecorder{}

	for _, v := range m {
		for i := range *v.CommitsJSON {
			if (*v.CommitsJSON)[i].Committer.Login == username {
				t := (*v.CommitsJSON)[i].Commit.Author.Date
				sr.AddCommit(t)
			}
		}
	}
	return sr.GetStreaks(), nil
}
