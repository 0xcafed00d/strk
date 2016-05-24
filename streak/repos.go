package streak

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func hasMorePages(h http.Header) bool {
	if link, ok := h["Link"]; ok {
		if strings.Contains(link[0], "rel=\"next\"") {
			return true
		}
	}
	return false
}

func GetRepos(user, apiToken string) (ReposJSON, error) {
	morePages := true
	var allrepos ReposJSON

	for pageNumber := 1; morePages; pageNumber++ {
		r, err := http.NewRequest("GET", makeGetReposURL(pageNumber, user), nil)
		if err != nil {
			return nil, err
		}
		r.Header["Authorization"] = []string{"token " + apiToken}

		client := &http.Client{}
		resp, err := client.Do(r)
		if err != nil {
			return nil, err
		}

		for k, v := range resp.Header {
			fmt.Println(k, v)
		}

		defer resp.Body.Close()

		var repos ReposJSON

		err = json.NewDecoder(resp.Body).Decode(&repos)
		if err != nil {
			return nil, err
		}

		allrepos = append(allrepos, repos...)
		fmt.Println(len(repos))

		morePages = hasMorePages(resp.Header)
	}
	fmt.Println(len(allrepos))

	return allrepos, nil
}
