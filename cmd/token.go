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
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	hiarc "github.com/hiarcdb/hiarc-go-sdk"
	"github.com/spf13/cobra"
)

var tokenExpires float32

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Token commands for Hiarc",
	Run:   nil,
}

var createUserTokenCmd = &cobra.Command{
	Use:   "create [user key]",
	Short: "Create a token scoped to a specific user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		tr := hiarc.CreateUserTokenRequest{Key: args[0]}

		if tokenExpires != 0 {
			tr.ExpirationMinues = tokenExpires
		}
		token, r, err := hiarcClient.TokenApi.CreateUserToken(context.Background(), tr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.CreateUserToken``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(token, "", "    ")
		fmt.Println(string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	tokenCmd.AddCommand(createUserTokenCmd)

	createUserTokenCmd.Flags().Float32Var(&tokenExpires, "expires-in", 0, "When token expires in minutes")
}
