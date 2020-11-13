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
	legalHoldName        string
	legalHoldDescription string
	legalHoldMetadata    string
)

var legalHoldCmd = &cobra.Command{
	Use:   "legal-hold",
	Short: "Legal Hold commands for Hiarc",
	Run:   nil,
}

var createLegalHoldCmd = &cobra.Command{
	Use:   "create [legal hold key]",
	Short: "Create Legal Hold with a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		clh := hiarc.CreateLegalHoldRequest{Key: args[0]}
		if legalHoldMetadata != "" {
			md, err := ConvertMetadataStringToObject(legalHoldMetadata)
			if err != nil {
				log.Fatal(err)
			}
			clh.Metadata = md
		}
		if legalHoldName != "" {
			clh.Name = legalHoldName
		}
		if legalHoldDescription != "" {
			clh.Description = legalHoldDescription
		}

		hold, r, err := hiarcClient.LegalHoldApi.CreateLegalHold(context.Background(), clh)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `LegalHoldApi.CreateLegalHold``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(hold, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getLegalHoldCmd = &cobra.Command{
	Use:   "get [legal hold key]",
	Short: "Get Legal Hold by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()

		hold, r, err := hiarcClient.LegalHoldApi.GetLegalHold(context.Background(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `LegalHoldApi.GetLegalHold``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(hold, "", "    ")
		fmt.Println(string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(legalHoldCmd)
	legalHoldCmd.AddCommand(createLegalHoldCmd)
	legalHoldCmd.AddCommand(getLegalHoldCmd)

	createLegalHoldCmd.Flags().StringVar(&legalHoldName, "name", "", "Legal Hold name")
	createLegalHoldCmd.Flags().StringVar(&legalHoldDescription, "description", "", "Legal Hold description")
	createLegalHoldCmd.Flags().StringVar(&legalHoldMetadata, "metadata", "", "Legal Hold metadata")
}
