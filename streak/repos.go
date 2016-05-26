package streak

import (
	"encoding/json"
	"fmt"
	"time"
)

type Repo struct {
	Name        string
	LastUpdated time.Time
	LastSHA     string
}

func makeGetReposURL(pageNumber int, user string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/repos?page=%d", user, pageNumber)
}

func GetRepos(user, apiToken string) (*ReposJSON, error) {
	var allrepos ReposJSON

	morePages := true
	for pageNumber := 1; morePages; pageNumber++ {
		resp, err := doAuthorizedRequest("GET", makeGetReposURL(pageNumber, user), apiToken, nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var repos ReposJSON
		err = json.NewDecoder(resp.Body).Decode(&repos)
		if err != nil {
			return nil, err
		}

		allrepos = append(allrepos, repos...)
		morePages = hasMorePages(resp.Header)
	}

	return &allrepos, nil
}
