package crawler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func requestAndParseJSON(url, body string, target interface{}) error {
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return err
	}
	defer request.Body.Close()
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	return json.NewDecoder(response.Body).Decode(target)
}
