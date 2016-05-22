package streak

import "time"

type Repo struct {
	Name        string
	LastUpdated time.Time
	LastSHA     string
}

func GetRepos(user string) ([]string, error) {

}
