package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	hiarc "github.com/allenmichael/hiarcgo"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

var (
	collectionMetadata    string
	collectionName        string
	collectionDescription string
	collectionQueries     []string
)

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Collection commands for Hiarc",
	Run:   nil,
}

var createCollectionCmd = &cobra.Command{
	Use:   "create [collection key]",
	Short: "Create a collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.CreateCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		ccr := hiarc.CreateCollectionRequest{Key: args[0]}
		if collectionMetadata != "" {
			md, err := ConvertMetadataStringToObject(collectionMetadata)
			if err != nil {
				log.Fatal(err)
			}
			ccr.Metadata = md
		}
		if collectionName != "" {
			ccr.Name = collectionName
		}
		if collectionDescription != "" {
			ccr.Description = collectionDescription
		}
		collection, r, err := hiarcClient.CollectionApi.CreateCollection(context.Background(), ccr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.CreateCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(collection, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getCollectionCmd = &cobra.Command{
	Use:   "get [collection key]",
	Short: "Get collection by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		collection, r, err := hiarcClient.CollectionApi.GetCollection(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(collection, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getAllCollectionsCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all collections",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetAllCollectionsOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		collections, r, err := hiarcClient.CollectionApi.GetAllCollections(context.Background(), &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetAllCollections``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(collections, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getChildrenForCollectionCmd = &cobra.Command{
	Use:   "children [collection key]",
	Short: "Get all child collections in a collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetCollectionChildrenOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		collection, r, err := hiarcClient.CollectionApi.GetCollectionChildren(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetCollectionChildren``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(collection, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getFilesForCollectionCmd = &cobra.Command{
	Use:   "files [collection key]",
	Short: "Get all files in a collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetCollectionFilesOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		collection, r, err := hiarcClient.CollectionApi.GetCollectionFiles(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetCollectionFiles``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(collection, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var getItemsForCollectionCmd = &cobra.Command{
	Use:   "items [collection key]",
	Short: "Get all files and child collections in a collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetCollectionItemsOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		collection, r, err := hiarcClient.CollectionApi.GetCollectionItems(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetCollectionFiles``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(collection, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var updateCollectionCmd = &cobra.Command{
	Use:   "update [collection key]",
	Short: "Update a collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.UpdateCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		ucr := hiarc.UpdateCollectionRequest{}
		if collectionMetadata != "" {
			md, err := ConvertMetadataStringToObject(collectionMetadata)
			if err != nil {
				log.Fatal(err)
			}
			ucr.Metadata = md
		}
		if collectionName != "" {
			ucr.Name = collectionName
		}
		if collectionDescription != "" {
			ucr.Description = collectionDescription
		}
		collection, r, err := hiarcClient.CollectionApi.UpdateCollection(context.Background(), args[0], ucr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.UpdateCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(collection, "", "    ")
		fmt.Println(string(jsonData))
	},
}

var deleteCollectionCmd = &cobra.Command{
	Use:   "delete [collection key]",
	Short: "Delete a collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.DeleteCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		_, r, err := hiarcClient.CollectionApi.DeleteCollection(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.DeleteCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			log.Fatal("Couldn't delete collection")
		}
		log.Println(fmt.Sprintf("Deleted collection: %s", args[0]))
	},
}

var removeFileFromCollectionCmd = &cobra.Command{
	Use:   "remove-file [collection key] [file key]",
	Short: "Remove a file from a collection",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.RemoveFileFromCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		_, r, err := hiarcClient.CollectionApi.RemoveFileFromCollection(context.Background(), args[0], args[1], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.RemoveFileFromCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			log.Fatal("Couldn't remove file from collection")
		}
		log.Println(fmt.Sprintf("Removed file %s from collection: %s", args[1], args[0]))
	},
}

var addUserToCollectionCmd = &cobra.Command{
	Use:   "add-user [collection key] [user key] [access level]",
	Short: "Add a user to a collection",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		accessLevel := strings.ToUpper(args[2])
		if !IsValidAccessLevel(accessLevel) {
			log.Println(fmt.Sprintf("%s is not a valid access level", accessLevel))
			log.Fatal(fmt.Sprintf("Choose from the following: %s, %s, %s, or %s", string(hiarc.CO_OWNER), string(hiarc.READ_WRITE), string(hiarc.READ_ONLY), string(hiarc.UPLOAD_ONLY)))
		}
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddUserToCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		al, err := GetAccessLevelFromString(accessLevel)
		if err != nil {
			log.Fatal(err)
		}
		aucr := hiarc.AddUserToCollectionRequest{UserKey: args[1], AccessLevel: al}
		_, r, err := hiarcClient.CollectionApi.AddUserToCollection(context.Background(), args[0], aucr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.AddUserToCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			log.Fatal("Couldn't add user to collection")
		}
		log.Println(fmt.Sprintf("Added user %s to collection %s with access level %s", args[1], args[0], accessLevel))
	},
}

var addGroupToCollectionCmd = &cobra.Command{
	Use:   "add-group [collection key] [group key] [access level]",
	Short: "Add a group to a collection",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		accessLevel := strings.ToUpper(args[2])
		if !IsValidAccessLevel(accessLevel) {
			log.Println(fmt.Sprintf("%s is not a valid access level", accessLevel))
			log.Fatal(fmt.Sprintf("Choose from the following: %s, %s, %s, or %s", string(hiarc.CO_OWNER), string(hiarc.READ_WRITE), string(hiarc.READ_ONLY), string(hiarc.UPLOAD_ONLY)))
		}
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddGroupToCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		al, err := GetAccessLevelFromString(accessLevel)
		if err != nil {
			log.Fatal(err)
		}
		agcr := hiarc.AddGroupToCollectionRequest{GroupKey: args[1], AccessLevel: al}
		_, r, err := hiarcClient.CollectionApi.AddGroupToCollection(context.Background(), args[0], agcr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.AddGroupToCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			log.Fatal("Couldn't add group to collection")
		}
		log.Println(fmt.Sprintf("Added group %s to collection %s with access level %s", args[1], args[0], accessLevel))
	},
}

var addFileToCollectionCmd = &cobra.Command{
	Use:   "add-file [collection key] [file key]",
	Short: "Add a file to a collection",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddFileToCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		afcr := hiarc.AddFileToCollectionRequest{FileKey: args[1]}
		_, r, err := hiarcClient.CollectionApi.AddFileToCollection(context.Background(), args[0], afcr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.AddFileToCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			log.Fatal("Couldn't add file to collection")
		}
		log.Println(fmt.Sprintf("Added file %s to collection %s", args[1], args[0]))
	},
}

var addChildToCollectionCmd = &cobra.Command{
	Use:   "add-child [parent collection key] [child collection key]",
	Short: "Add a child to a collection",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddChildToCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}
		_, r, err := hiarcClient.CollectionApi.AddChildToCollection(context.Background(), args[0], args[1], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.AddChildToCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			log.Fatal("Couldn't add child to collection")
		}
		log.Println(fmt.Sprintf("Added child %s to collection %s", args[1], args[0]))
	},
}

var findCollectionCmd = &cobra.Command{
	Use:   "find",
	Short: "Find collection by query",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.FindCollectionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		queries := make([]map[string]interface{}, 0)
		for i := range collectionQueries {
			q, err := ConvertQueryToObject(collectionQueries[i])
			if err != nil {
				log.Fatal(err)
			}
			queries = append(queries, q)
		}
		qr := hiarc.FindCollectionsRequest{Query: queries}
		fc, r, err := hiarcClient.CollectionApi.FindCollection(context.Background(), qr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.FindCollection``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(fc, "", "    ")
		fmt.Println(string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(createCollectionCmd)
	collectionCmd.AddCommand(getCollectionCmd)
	collectionCmd.AddCommand(updateCollectionCmd)
	collectionCmd.AddCommand(deleteCollectionCmd)
	collectionCmd.AddCommand(removeFileFromCollectionCmd)
	collectionCmd.AddCommand(addGroupToCollectionCmd)
	collectionCmd.AddCommand(addUserToCollectionCmd)
	collectionCmd.AddCommand(addFileToCollectionCmd)
	collectionCmd.AddCommand(addChildToCollectionCmd)
	collectionCmd.AddCommand(findCollectionCmd)

	getCollectionCmd.AddCommand(getAllCollectionsCmd)
	getCollectionCmd.AddCommand(getChildrenForCollectionCmd)
	getCollectionCmd.AddCommand(getItemsForCollectionCmd)
	getCollectionCmd.AddCommand(getFilesForCollectionCmd)

	createCollectionCmd.Flags().StringVar(&collectionName, "name", "", "Collection name")
	createCollectionCmd.Flags().StringVar(&collectionDescription, "description", "", "Collection description")
	createCollectionCmd.Flags().StringVar(&collectionMetadata, "metadata", "", "Collection metadata")

	updateCollectionCmd.Flags().StringVar(&collectionName, "name", "", "Collection name")
	updateCollectionCmd.Flags().StringVar(&collectionDescription, "description", "", "Collection description")
	updateCollectionCmd.Flags().StringVar(&collectionMetadata, "metadata", "", "Collection metadata")

	findCollectionCmd.Flags().StringArrayVar(&collectionQueries, "query", make([]string, 0), "Collection query")
	findCollectionCmd.MarkFlagRequired("query")
}
