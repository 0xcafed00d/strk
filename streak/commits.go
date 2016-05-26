package streak

import (
	"encoding/json"
	"fmt"
)

func makeGetCommitsURL(pageNumber int, user, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?page=%d", user, repo, pageNumber)
}

func GetCommits(user, repo, apiToken string) (*CommitsJSON, error) {
	var allcommits CommitsJSON

	morePages := true
	for pageNumber := 1; morePages; pageNumber++ {
		resp, err := doAuthorizedRequest("GET", makeGetCommitsURL(pageNumber, user, repo), apiToken, nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var commits CommitsJSON
		err = json.NewDecoder(resp.Body).Decode(&commits)
		if err != nil {
			return nil, err
		}

		allcommits = append(allcommits, commits...)
		morePages = hasMorePages(resp.Header)
	}

	return &allcommits, nil
}
