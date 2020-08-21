package hiarcapi

import (
	"context"
	"fmt"
	"os"

	openapi "github.com/hiarcdb/hiarc-go-sdk"
)

func GetUserOpenApi(key string) (*openapi.User, error) {

	configuration := openapi.NewConfiguration()
	// configuration.Servers = openapi.ServerConfigurations{
	// 	{
	// 		URL:         "http://localhost:5000",
	// 		Description: "Localhost",
	// 	},
	// }
	configuration.AddDefaultHeader("X-Hiarc-Api-Key", "adminkey")
	configuration.AddDefaultHeader("Content-type", "application/json")

	apiClient := openapi.NewAPIClient(configuration)
	user, r, err := apiClient.UserApi.GetUser(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return nil, err
	}

	return &user, err
}

// This was the old way before openapi sdk
//https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
//  var client = http.Client{}

//  func GetUser(key string) (map[string]interface{}, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:5000/users/%s", key), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("X-Hiarc-Api-Key", "adminkey")
// 	req.Header.Set("Content-type", "application/json")
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	var result map[string]interface{}
// 	json.NewDecoder(resp.Body).Decode(&result)

// 	return result, err
// }
