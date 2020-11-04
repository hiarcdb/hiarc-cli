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
	"log"
	"os"

	hiarc "github.com/allenmichael/hiarcgo"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

var (
	key             string
	userName        string
	userDescription string
	userMetadata    string
	userKey         string
	userQueries     []string
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User commands for Hiarc",
	Run:   nil,
}

var createUserCmd = &cobra.Command{
	Use:   "create [user key]",
	Short: "Create user with a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		cu := hiarc.CreateUserRequest{Key: args[0]}
		if userMetadata != "" {
			md, err := ConvertMetadataStringToObject(userMetadata)
			if err != nil {
				log.Fatal(err)
			}
			cu.Metadata = md
		}
		if userName != "" {
			cu.Name = userName
		}
		if userDescription != "" {
			cu.Description = userDescription
		}

		user, r, err := hiarcClient.UserApi.CreateUser(context.Background(), cu)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(user, "", "    ")
		log.Println(string(jsonData))
	},
}

var getUserCmd = &cobra.Command{
	Use:   "get [user key]",
	Short: "Get user by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		user, r, err := hiarcClient.UserApi.GetUser(context.Background(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(user, "", "    ")
		log.Println(string(jsonData))
	},
}

var getAllUsersCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all users",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		user, r, err := hiarcClient.UserApi.GetAllUsers(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(user, "", "    ")
		log.Println(string(jsonData))
	},
}

var updateUserCmd = &cobra.Command{
	Use:   "update [user key]",
	Short: "Update user by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		uu := hiarc.UpdateUserRequest{}
		if userMetadata != "" {
			md, err := ConvertMetadataStringToObject(userMetadata)
			if err != nil {
				log.Fatal(err)
			}
			uu.Metadata = md
		}
		if userName != "" {
			uu.Name = userName
		}
		if userDescription != "" {
			uu.Description = userDescription
		}
		user, r, err := hiarcClient.UserApi.UpdateUser(context.Background(), args[0], uu)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Update``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(user, "", "    ")
		log.Println(string(jsonData))
	},
}

var deleteUserCmd = &cobra.Command{
	Use:   "delete [user key]",
	Short: "Delete user by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		_, r, err := hiarcClient.UserApi.DeleteUser(context.Background(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		log.Println(fmt.Sprintf("Deleted user: %s", args[0]))
	},
}

var findUserCmd = &cobra.Command{
	Use:   "find",
	Short: "Find user by query",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		queries := make([]map[string]interface{}, 0)
		for i := range userQueries {
			q, err := ConvertQueryToObject(userQueries[i])
			if err != nil {
				log.Fatal(err)
			}
			queries = append(queries, q)
		}
		qr := hiarc.FindUsersRequest{Query: queries}
		fu, r, err := hiarcClient.UserApi.FindUser(context.Background(), qr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Update``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(fu, "", "    ")
		log.Println(string(jsonData))
	},
}

var getCurrentUserCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the current",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetCurrentUserOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		user, r, err := hiarcClient.UserApi.GetCurrentUser(context.Background(), &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(user, "", "    ")
		log.Println(string(jsonData))
	},
}

var getGroupsForUserCmd = &cobra.Command{
	Use:   "groups",
	Short: "Get groups for a user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Needs user key as argument")
		}
		hiarcClient := ConfigureHiarcClient()

		groups, r, err := hiarcClient.UserApi.GetGroupsForUser(context.Background(), args[0], nil)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(groups, "", "    ")
		log.Println(string(jsonData))
	},
}

var getGroupsForCurrentUserCmd = &cobra.Command{
	Use:   "groups",
	Short: "Get groups for current user",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetGroupsForCurrentUserOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		groups, r, err := hiarcClient.UserApi.GetGroupsForCurrentUser(context.Background(), &opts)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(groups, "", "    ")
		log.Println(string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(createUserCmd)
	userCmd.AddCommand(getUserCmd)
	userCmd.AddCommand(updateUserCmd)
	userCmd.AddCommand(deleteUserCmd)
	userCmd.AddCommand(findUserCmd)
	getUserCmd.AddCommand(getAllUsersCmd)
	getUserCmd.AddCommand(getCurrentUserCmd)
	getUserCmd.AddCommand(getGroupsForUserCmd)
	getCurrentUserCmd.AddCommand(getGroupsForCurrentUserCmd)

	createUserCmd.Flags().StringVar(&userName, "name", "", "User name")
	createUserCmd.Flags().StringVar(&userDescription, "description", "", "User description")
	createUserCmd.Flags().StringVar(&userMetadata, "metadata", "", "User metadata")

	updateUserCmd.Flags().StringVar(&userName, "name", "", "User name")
	updateUserCmd.Flags().StringVar(&userDescription, "description", "", "User description")
	updateUserCmd.Flags().StringVar(&userMetadata, "metadata", "", "User metadata")

	findUserCmd.Flags().StringArrayVar(&userQueries, "query", make([]string, 0), "User query")
	findUserCmd.MarkFlagRequired("query")
}
