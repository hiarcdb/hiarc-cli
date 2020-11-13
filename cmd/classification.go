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
	classificationName        string
	classificationDescription string
	classificationMetadata    string
	classificationQueries     []string
)

var classificationCmd = &cobra.Command{
	Use:   "classification",
	Short: "Classification commands for Hiarc",
	Run:   nil,
}

var createClassificationCmd = &cobra.Command{
	Use:   "create [classification key]",
	Short: "Create classification with a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.CreateClassificationOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		ccr := hiarc.CreateClassificationRequest{Key: args[0]}
		if classificationMetadata != "" {
			md, err := ConvertMetadataStringToObject(classificationMetadata)
			if err != nil {
				log.Fatal(err)
			}
			ccr.Metadata = md
		}
		if classificationName != "" {
			ccr.Name = classificationName
		}
		if classificationDescription != "" {
			ccr.Description = classificationDescription
		}

		cc, r, err := hiarcClient.ClassificationApi.CreateClassification(context.Background(), ccr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ClassificationApi.CreateClassification``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(cc, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getClassificationCmd = &cobra.Command{
	Use:   "get [classification key]",
	Short: "Get classification by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetClassificationOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		classification, r, err := hiarcClient.ClassificationApi.GetClassification(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ClassificationApi.GetClassification``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(classification, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getAllClassificationsCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all classifications",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetAllClassificationsOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		classifications, r, err := hiarcClient.ClassificationApi.GetAllClassifications(context.Background(), &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ClassificationApi.GetAllClassifications``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(classifications, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var updateClassificationCmd = &cobra.Command{
	Use:   "update [classification key]",
	Short: "Update classification by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.UpdateClassificationOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		uc := hiarc.UpdateClassificationRequest{}
		if classificationMetadata != "" {
			md, err := ConvertMetadataStringToObject(classificationMetadata)
			if err != nil {
				log.Fatal(err)
			}
			uc.Metadata = md
		}
		if classificationName != "" {
			uc.Name = classificationName
		}
		if classificationDescription != "" {
			uc.Description = classificationDescription
		}
		classification, r, err := hiarcClient.ClassificationApi.UpdateClassification(context.Background(), args[0], uc, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ClassificationApi.UpdateClassification``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(classification, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var deleteClassificationCmd = &cobra.Command{
	Use:   "delete [classification key]",
	Short: "Delete classification by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.DeleteClassificationOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		_, r, err := hiarcClient.ClassificationApi.DeleteClassification(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ClassificationApi.DeleteClassification``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		log.Println(fmt.Sprintf("Deleted classification: %s", args[0]))
	},
}

var findClassificationCmd = &cobra.Command{
	Use:   "find",
	Short: "Find classification by query",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.FindClassificationOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		queries := make([]map[string]interface{}, 0)
		for i := range classificationQueries {
			q, err := ConvertQueryToObject(classificationQueries[i])
			if err != nil {
				log.Fatal(err)
			}
			queries = append(queries, q)
		}
		qr := hiarc.FindClassificationsRequest{Query: queries}
		fc, r, err := hiarcClient.ClassificationApi.FindClassification(context.Background(), qr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ClassificationApi.FindClassification``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(fc, "", "    ")
		fmt.Println(string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(classificationCmd)
	classificationCmd.AddCommand(createClassificationCmd)
	classificationCmd.AddCommand(getClassificationCmd)
	classificationCmd.AddCommand(updateClassificationCmd)
	// classificationCmd.AddCommand(deleteClassificationCmd)
	classificationCmd.AddCommand(findClassificationCmd)

	getClassificationCmd.AddCommand(getAllClassificationsCmd)

	createClassificationCmd.Flags().StringVar(&classificationName, "name", "", "Classification name")
	createClassificationCmd.Flags().StringVar(&classificationDescription, "description", "", "Classification description")
	createClassificationCmd.Flags().StringVar(&classificationMetadata, "metadata", "", "Classification metadata")

	updateClassificationCmd.Flags().StringVar(&classificationName, "name", "", "Classification name")
	updateClassificationCmd.Flags().StringVar(&classificationDescription, "description", "", "Classification description")
	updateClassificationCmd.Flags().StringVar(&classificationMetadata, "metadata", "", "Classification metadata")

	findClassificationCmd.Flags().StringArrayVar(&classificationQueries, "query", make([]string, 0), "Classification query")
	findClassificationCmd.MarkFlagRequired("query")
}
