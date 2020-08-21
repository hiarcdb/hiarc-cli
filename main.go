/*
Copyright © 2020 Hiarc <support@hiarcdb.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hiarcdb/hiarc-cli/cmd"
	openapiclient "github.com/hiarcdb/hiarc-go-sdk"
)

func main() {
	key := "user-1" // string | Key of user to get

	configuration := openapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-Hiarc-Api-Key", "adminkey")
	configuration.AddDefaultHeader("Content-type", "application/json")

	apiClient := openapiclient.NewAPIClient(configuration)
	user, r, err := apiClient.UserApi.GetUser(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	jsonData, err := json.MarshalIndent(user, "", "    ")
	log.Println(string(jsonData))

	cmd.Execute()
}
