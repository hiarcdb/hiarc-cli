package hiarcapi

import (
	"encoding/json"
	"net/http"
)

//https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7

var client = http.Client{}

func GetUser() (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", "http://localhost:5000/users/user-1", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Hiarc-Api-Key", "adminkey")
	req.Header.Set("Content-type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, err
}
