package streak

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Repo struct {
	Name        string
	LastUpdated time.Time
	LastSHA     string
}

func GetRepos(user string) ([]string, error) {

	r, err := http.NewRequest("GET", "https://api.github.com/users/"+user+"/repos?page=1&per_page=10", nil)
	if err != nil {
		return nil, err
	}

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

	fmt.Println(len(repos))

	return nil, nil
}
