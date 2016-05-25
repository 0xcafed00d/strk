package streak

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func hasMorePages(h http.Header) bool {
	if link, ok := h["Link"]; ok {
		if strings.Contains(link[0], "rel=\"next\"") {
			return true
		}
	}
	return false
}

func statusAsError(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP response: %s", resp.Status)
	}
	return nil
}

func doAuthorizedRequest(method, URL, apiToken string, body io.Reader) (resp *http.Response, err error) {
	var r *http.Request
	r, err = http.NewRequest(method, URL, body)
	if err != nil {
		return
	}
	r.Header["Authorization"] = []string{"token " + apiToken}

	client := &http.Client{}
	resp, err = client.Do(r)
	if err != nil {
		return
	}
	err = statusAsError(resp)
	return
}
