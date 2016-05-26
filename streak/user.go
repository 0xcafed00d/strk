package streak

import (
	"bytes"
	"encoding/json"
)

func GetUser(apiToken string) (*UserJSON, error) {

	resp, err := doAuthorizedRequest("GET", "https://api.github.com/user", apiToken, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user UserJSON
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func UpdateUser(updateStruct interface{}, apiToken string) error {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(updateStruct)

	resp, err := doAuthorizedRequest("PATCH", "https://api.github.com/user", apiToken, &buf)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return err
}
