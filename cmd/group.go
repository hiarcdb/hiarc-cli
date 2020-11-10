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
	groupMetadata    string
	groupName        string
	groupDescription string
	groupQueries     []string
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Group commands for Hiarc",
	Run:   nil,
}

var createGroupCmd = &cobra.Command{
	Use:   "create [group key]",
	Short: "Create a group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		cgr := hiarc.CreateGroupRequest{Key: args[0]}
		if groupMetadata != "" {
			md, err := ConvertMetadataStringToObject(groupMetadata)
			if err != nil {
				log.Fatal(err)
			}
			cgr.Metadata = md
		}
		if groupName != "" {
			cgr.Name = groupName
		}
		if groupDescription != "" {
			cgr.Description = groupDescription
		}
		group, r, err := hiarcClient.GroupApi.CreateGroup(context.Background(), cgr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.CreateGroup``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(group, "", "    ")
		log.Println(string(jsonData))
	},
}

var getGroupCmd = &cobra.Command{
	Use:   "get [group key]",
	Short: "Get group by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		group, r, err := hiarcClient.GroupApi.GetGroup(context.Background(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroup``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(group, "", "    ")
		log.Println(string(jsonData))
	},
}

var getAllGroupsCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all groups",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		groups, r, err := hiarcClient.GroupApi.GetAllGroups(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetAllGroups``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(groups, "", "    ")
		log.Println(string(jsonData))
	},
}

var getGroupsForUserGroupCmd = &cobra.Command{
	Use:   "for-user [user key]",
	Short: "Get all groups for a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetGroupsForUserOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		group, r, err := hiarcClient.GroupsApi.GetGroupsForUser(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroupsForCurrentUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(group, "", "    ")
		log.Println(string(jsonData))
	},
}

var getGroupsCurrentUserCmd = &cobra.Command{
	Use:   "current",
	Short: "Get all groups for current user",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetGroupsForCurrentUserOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		group, r, err := hiarcClient.GroupApi.GetGroupsForCurrentUser(context.Background(), &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroupsForCurrentUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(group, "", "    ")
		log.Println(string(jsonData))
	},
}

var updateGroupCmd = &cobra.Command{
	Use:   "update [group key]",
	Short: "Update a group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		ugr := hiarc.UpdateGroupRequest{}
		if groupMetadata != "" {
			md, err := ConvertMetadataStringToObject(groupMetadata)
			if err != nil {
				log.Fatal(err)
			}
			ugr.Metadata = md
		}
		if groupName != "" {
			ugr.Name = groupName
		}
		if groupDescription != "" {
			ugr.Description = groupDescription
		}
		group, r, err := hiarcClient.GroupApi.UpdateGroup(context.Background(), args[0], ugr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.UpdateGroup``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(group, "", "    ")
		log.Println(string(jsonData))
	},
}

var deleteGroupCmd = &cobra.Command{
	Use:   "delete [group key]",
	Short: "Delete a group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		_, r, err := hiarcClient.GroupApi.DeleteGroup(context.Background(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteGroup``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		log.Println(fmt.Sprintf("Deleted group: %s", args[0]))
	},
}

var addUserToGroupCmd = &cobra.Command{
	Use:   "add-user [group key] [user key]",
	Short: "Add a user to a group",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		_, r, err := hiarcClient.GroupApi.AddUserToGroup(context.Background(), args[0], args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.AddUserToGroup``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		log.Println(fmt.Sprintf("Added user %s to group %s", args[1], args[0]))
	},
}

var findGroupCmd = &cobra.Command{
	Use:   "find",
	Short: "Find group by query",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		queries := make([]map[string]interface{}, 0)
		for i := range groupQueries {
			q, err := ConvertQueryToObject(groupQueries[i])
			if err != nil {
				log.Fatal(err)
			}
			queries = append(queries, q)
		}
		qr := hiarc.FindGroupsRequest{Query: queries}
		fg, r, err := hiarcClient.GroupApi.FindGroup(context.Background(), qr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.FindGroup``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(fg, "", "    ")
		log.Println(string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(createGroupCmd)
	groupCmd.AddCommand(getGroupCmd)
	groupCmd.AddCommand(updateGroupCmd)
	groupCmd.AddCommand(deleteGroupCmd)
	groupCmd.AddCommand(addUserToGroupCmd)
	groupCmd.AddCommand(findGroupCmd)

	getGroupCmd.AddCommand(getAllGroupsCmd)
	getGroupCmd.AddCommand(getGroupsForUserGroupCmd)

	getAllGroupsCmd.AddCommand(getGroupsCurrentUserCmd)

	createGroupCmd.Flags().StringVar(&groupName, "name", "", "Group name")
	createGroupCmd.Flags().StringVar(&groupDescription, "description", "", "Group description")
	createGroupCmd.Flags().StringVar(&groupMetadata, "metadata", "", "Group metadata")

	updateGroupCmd.Flags().StringVar(&groupName, "name", "", "Group name")
	updateGroupCmd.Flags().StringVar(&groupDescription, "description", "", "Group description")
	updateGroupCmd.Flags().StringVar(&groupMetadata, "metadata", "", "Group metadata")

	findGroupCmd.Flags().StringArrayVar(&groupQueries, "query", make([]string, 0), "Group query")
	findGroupCmd.MarkFlagRequired("query")
}
