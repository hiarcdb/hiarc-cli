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
	"github.com/spf13/cobra"
)

var (
	retentionName        string
	retentionDescription string
	retentionMetadata    string
	retentionQueries     []string
)

var retentionCmd = &cobra.Command{
	Use:   "retention-policy",
	Short: "Retention Policy commands for Hiarc",
	Run:   nil,
}

var createRetentionCmd = &cobra.Command{
	Use:   "create [retention policy key]",
	Short: "Create Retention Policy with a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		cr := hiarc.CreateRetentionPolicyRequest{Key: args[0]}
		if retentionMetadata != "" {
			md, err := ConvertMetadataStringToObject(retentionMetadata)
			if err != nil {
				log.Fatal(err)
			}
			cr.Metadata = md
		}
		if retentionName != "" {
			cr.Name = retentionName
		}
		if retentionDescription != "" {
			cr.Description = retentionDescription
		}

		policy, r, err := hiarcClient.RetentionPolicyApi.CreateRetentionPolicy(context.Background(), cr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `RetentionPolicyApi.CreateRetentionPolicy``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(policy, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getRetentionCmd = &cobra.Command{
	Use:   "get [retention policy key]",
	Short: "Get Retention Policy by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		retention, r, err := hiarcClient.RetentionPolicyApi.GetRetentionPolicy(context.Background(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `RetentionPolicyApi.GetRetentionPolicy``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(retention, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getAllPoliciesCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all Retention Policies",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		policies, r, err := hiarcClient.RetentionPolicyApi.GetAllRetentionPolicies(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `RetentionPolicyApi.GetAllRetentionPolicies``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(policies, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var updateRetentionCmd = &cobra.Command{
	Use:   "update [retention policy key]",
	Short: "Update Retention Policy by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		ur := hiarc.UpdateRetentionPolicyRequest{}
		if retentionMetadata != "" {
			md, err := ConvertMetadataStringToObject(retentionMetadata)
			if err != nil {
				log.Fatal(err)
			}
			ur.Metadata = md
		}
		if retentionName != "" {
			ur.Name = retentionName
		}
		if retentionDescription != "" {
			ur.Description = retentionDescription
		}
		retention, r, err := hiarcClient.RetentionPolicyApi.UpdateRetentionPolicy(context.Background(), args[0], ur)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `RetentionPolicyApi.UpdateRetentionPolicy``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(retention, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var findRetentionCmd = &cobra.Command{
	Use:   "find",
	Short: "Find Retention Policy by query",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		queries := make([]map[string]interface{}, 0)
		for i := range retentionQueries {
			q, err := ConvertQueryToObject(retentionQueries[i])
			if err != nil {
				log.Fatal(err)
			}
			queries = append(queries, q)
		}
		qr := hiarc.FindRetentionPoliciesRequest{Query: queries}
		fr, r, err := hiarcClient.RetentionPolicyApi.FindRetentionPolicies(context.Background(), qr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `RetentionPolicyApi.FindRetentionPolicies``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(fr, "", "    ")
		fmt.Println(string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(retentionCmd)
	retentionCmd.AddCommand(createRetentionCmd)
	retentionCmd.AddCommand(getRetentionCmd)
	retentionCmd.AddCommand(updateRetentionCmd)
	retentionCmd.AddCommand(findRetentionCmd)
	getRetentionCmd.AddCommand(getAllPoliciesCmd)

	createRetentionCmd.Flags().StringVar(&retentionName, "name", "", "Retention name")
	createRetentionCmd.Flags().StringVar(&retentionDescription, "description", "", "Retention description")
	createRetentionCmd.Flags().StringVar(&retentionMetadata, "metadata", "", "Retention metadata")

	updateRetentionCmd.Flags().StringVar(&retentionName, "name", "", "Retention name")
	updateRetentionCmd.Flags().StringVar(&retentionDescription, "description", "", "Retention description")
	updateRetentionCmd.Flags().StringVar(&retentionMetadata, "metadata", "", "Retention metadata")

	findRetentionCmd.Flags().StringArrayVar(&retentionQueries, "query", make([]string, 0), "Retention query")
	findRetentionCmd.MarkFlagRequired("query")
}
