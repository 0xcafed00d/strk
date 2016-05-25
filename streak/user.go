package streak

func GetUser(apiToken string) (*int, error) {

	resp, err := doAuthorizedRequest("GET", "https://api.github.com/user", apiToken, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return nil, err
}
