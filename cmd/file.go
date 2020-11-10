package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	hiarc "github.com/allenmichael/hiarcgo"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

var (
	fileName            string
	fileDescription     string
	fileMetadata        string
	fileKey             string
	fileAccessLevel     string
	fileStorageService  string
	fileStorageId       string
	filePathUpload      string
	filePathDownload    string
	directUploadExpires int32
	fileQueries         []string
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "File commands for Hiarc",
	Run:   nil,
}

var getFileCmd = &cobra.Command{
	Use:   "get [file key]",
	Short: "Get file by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		file, r, err := hiarcClient.FileApi.GetFile(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var getFileVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Get versions of file by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Needs file key as argument")
		}
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetVersionsOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		versions, r, err := hiarcClient.FileApi.GetVersions(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetVersions``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(versions, "", "    ")
		log.Println(string(jsonData))
	},
}

var getFileRetentionPoliciesCmd = &cobra.Command{
	Use:   "retention-policies",
	Short: "Get retention policies for file by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Needs file key as argument")
		}
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetRetentionPoliciesOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		policies, r, err := hiarcClient.FileApi.GetRetentionPolicies(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetRetentionPolicies``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(policies, "", "    ")
		log.Println(string(jsonData))
	},
}

var getFileCollectionsCmd = &cobra.Command{
	Use:   "collections",
	Short: "Get collections for file by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Needs file key as argument")
		}
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetCollectionsForFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		collections, r, err := hiarcClient.FileApi.GetCollectionsForFile(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetCollectionsForFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(collections, "", "    ")
		log.Println(string(jsonData))
	},
}

var getDirectDownloadCmd = &cobra.Command{
	Use:   "direct-download [file key]",
	Short: "Get direct download for file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.GetDirectDownloadUrlOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		url, r, err := hiarcClient.FileApi.GetDirectDownloadUrl(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetRetentionPolicies``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(url, "", "    ")
		log.Println(string(jsonData))
	},
}

var getDirectUploadCmd = &cobra.Command{
	Use:   "direct-upload",
	Short: "Get direct upload for file",
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.CreateDirectUploadUrlOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}
		if directUploadExpires != 0 {
			opts.ExpiresInSeconds = optional.NewInt32(directUploadExpires)
		}

		du := hiarc.CreateDirectUploadUrlRequest{}
		if fileStorageService != "" {
			du.StorageService = fileStorageService
		}

		url, r, err := hiarcClient.FileApi.CreateDirectUploadUrl(context.Background(), du, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.CreateDirectUploadUrl``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(url, "", "    ")
		log.Println(string(jsonData))
	},
}

var createFileCmd = &cobra.Command{
	Use:   "create [file key]",
	Short: "Upload a file with key and other file attributes",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.CreateFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		cf := hiarc.CreateFileRequest{Key: args[0]}
		if fileMetadata != "" {
			md, err := ConvertMetadataStringToObject(fileMetadata)
			if err != nil {
				log.Fatal(err)
			}
			cf.Metadata = md
		}
		if fileName != "" {
			cf.Name = fileName
		} else {
			fi, err := os.Stat(filePathUpload)
			if err != nil {
				log.Fatal(err)
			}
			fileName = fi.Name()
			cf.Name = fi.Name()
		}
		if fileDescription != "" {
			cf.Description = fileDescription
		}
		if fileStorageService != "" {
			cf.StorageService = fileStorageService
		}

		file, r, err := hiarcClient.FileApi.CreateFile(context.Background(), filePathUpload, fileName, cf, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.CreateFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var attachFileCmd = &cobra.Command{
	Use:   "attach [file key]",
	Short: "Attach to an existing file in a storage service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AttachToExisitingFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		ar := hiarc.AttachToExistingFileRequest{}
		if fileName != "" {
			ar.Name = fileName
		}
		if fileStorageService != "" {
			ar.StorageService = fileStorageService
		}
		if fileStorageId != "" {
			ar.StorageId = fileStorageId
		}

		file, r, err := hiarcClient.FileApi.AttachToExisitingFile(context.Background(), args[0], ar, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.AttachToExisitingFIle``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var copyFileCmd = &cobra.Command{
	Use:   "copy [source file key] [destination file key]",
	Short: "Attach to an existing file in a storage service",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.CopyFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		cr := hiarc.CopyFileRequest{Key: args[1]}
		if fileStorageService != "" {
			cr.StorageService = fileStorageService
		}

		file, r, err := hiarcClient.FileApi.CopyFile(context.Background(), args[0], cr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.CopyFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var addVersionCmd = &cobra.Command{
	Use:   "add-version [file key]",
	Short: "Upload a new version of a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddVersionOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		av := hiarc.AddVersionToFileRequest{Key: args[0]}
		if fileStorageService != "" {
			av.StorageService = fileStorageService
		}
		if fileName == "" {
			fi, err := os.Stat(filePathUpload)
			if err != nil {
				log.Fatal(err)
			}
			fileName = fi.Name()
		}

		file, r, err := hiarcClient.FileApi.AddVersion(context.Background(), args[0], filePathUpload, fileName, av, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.AddVersion``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var addGroupToFileCmd = &cobra.Command{
	Use:   "add-group [file key] [group key] [access level]",
	Short: "Grant access to a file for a group",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		accessLevel := strings.ToUpper(args[2])
		if !IsValidAccessLevel(accessLevel) {
			log.Println(fmt.Sprintf("%s is not a valid access level", accessLevel))
			log.Fatal(fmt.Sprintf("Choose from the following: %s, %s, %s, or %s", string(hiarc.CO_OWNER), string(hiarc.READ_WRITE), string(hiarc.READ_ONLY), string(hiarc.UPLOAD_ONLY)))
		}
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddGroupToFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}
		al, err := GetAccessLevelFromString(accessLevel)
		if err != nil {
			log.Fatal(err)
		}
		ag := hiarc.AddGroupToFileRequest{GroupKey: args[1], AccessLevel: al}

		file, r, err := hiarcClient.FileApi.AddGroupToFile(context.Background(), args[0], ag, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.AddGroupToFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var addUserToFileCmd = &cobra.Command{
	Use:   "add-user [file key] [user key] [access level]",
	Short: "Grant access to a file for a user",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		accessLevel := strings.ToUpper(args[2])
		if !IsValidAccessLevel(accessLevel) {
			log.Println(fmt.Sprintf("%s is not a valid access level", accessLevel))
			log.Fatal(fmt.Sprintf("Choose from the following: %s, %s, %s, or %s", string(hiarc.CO_OWNER), string(hiarc.READ_WRITE), string(hiarc.READ_ONLY), string(hiarc.UPLOAD_ONLY)))
		}
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddUserToFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}
		al, err := GetAccessLevelFromString(accessLevel)
		if err != nil {
			log.Fatal(err)
		}
		au := hiarc.AddUserToFileRequest{UserKey: args[1], AccessLevel: al}

		file, r, err := hiarcClient.FileApi.AddUserToFile(context.Background(), args[0], au, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.AddUserToFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var addClassificationToFileCmd = &cobra.Command{
	Use:   "add-classification [file key] [classification key]",
	Short: "Add classification to a file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddClassificationToFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}
		ac := hiarc.AddClassificationToFileRequest{ClassificationKey: args[1]}

		file, r, err := hiarcClient.FileApi.AddClassificationToFile(context.Background(), args[0], ac, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.AddClassificationToFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var addRetentionPolicyToFileCmd = &cobra.Command{
	Use:   "add-retention [file key] [retention policy key]",
	Short: "Add a retention policy to a file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.AddRetentionPolicyToFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}
		ar := hiarc.AddRetentionPolicyToFileRequest{RetentionPolicyKey: args[1]}

		file, r, err := hiarcClient.FileApi.AddRetentionPolicyToFile(context.Background(), args[0], ar, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.AddRetentionPolicyToFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var downloadFileCmd = &cobra.Command{
	Use:   "download [file key]",
	Short: "Download a file to your local system",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.DownloadFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}
		getOpts := hiarc.GetFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}
		s, err := os.Stat(filePathDownload)
		if err != nil {
			log.Fatal(err)
		}
		if s.IsDir() != true {
			log.Fatal("Download path must be a directory.")
		}
		if fileName == "" {
			f, r, err := hiarcClient.FileApi.GetFile(context.Background(), args[0], &getOpts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `FileApi.GetFile``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			}
			fileName = f.Name
		}

		fib, r, err := hiarcClient.FileApi.DownloadFile(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.DownloadFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}

		out, err := os.Create(filepath.Join(filePathDownload, fileName))
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(out, fib)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(fmt.Sprintf("Downloaded file: %s to the following location: %s", args[0], filePathDownload))
	},
}

var updateFileCmd = &cobra.Command{
	Use:   "update [file key]",
	Short: "Update a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.UpdateFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		uf := hiarc.UpdateFileRequest{}
		if fileMetadata != "" {
			md, err := ConvertMetadataStringToObject(fileMetadata)
			if err != nil {
				log.Fatal(err)
			}
			uf.Metadata = md
		}
		if fileName != "" {
			uf.Name = fileName
		}
		if fileDescription != "" {
			uf.Description = fileDescription
		}

		file, r, err := hiarcClient.FileApi.UpdateFile(context.Background(), args[0], uf, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.UpdateFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var filterFilesCmd = &cobra.Command{
	Use:   "filter [list of file keys]",
	Short: "Filter which files a user can access",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.FilterAllowedFilesOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		fr := hiarc.AllowedFilesRequest{Keys: args}

		file, r, err := hiarcClient.FilesApi.FilterAllowedFiles(context.Background(), fr, &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.FilterAllowedFiles``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		jsonData, err := json.MarshalIndent(file, "", "    ")
		log.Println(string(jsonData))
	},
}

var deleteFileCmd = &cobra.Command{
	Use:   "delete [file key]",
	Short: "Delete file by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hiarcClient := ConfigureHiarcClient()
		asUser, err := rootCmd.Flags().GetString("as-user")
		opts := hiarc.DeleteFileOpts{}
		if asUser != "" && err == nil {
			opts.XHiarcUserKey = optional.NewString(asUser)
		}

		_, r, err := hiarcClient.FileApi.DeleteFile(context.Background(), args[0], &opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FileApi.DeleteFile``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		} else {
			log.Println(fmt.Sprintf("Deleted file: %s", args[0]))
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	fileCmd.AddCommand(getFileCmd)
	fileCmd.AddCommand(createFileCmd)
	fileCmd.AddCommand(attachFileCmd)
	fileCmd.AddCommand(updateFileCmd)
	fileCmd.AddCommand(downloadFileCmd)
	fileCmd.AddCommand(deleteFileCmd)
	fileCmd.AddCommand(addVersionCmd)
	fileCmd.AddCommand(addUserToFileCmd)
	fileCmd.AddCommand(addGroupToFileCmd)
	fileCmd.AddCommand(addRetentionPolicyToFileCmd)
	fileCmd.AddCommand(addClassificationToFileCmd)
	fileCmd.AddCommand(getDirectDownloadCmd)
	fileCmd.AddCommand(getDirectUploadCmd)
	fileCmd.AddCommand(copyFileCmd)
	fileCmd.AddCommand(filterFilesCmd)

	getFileCmd.AddCommand(getFileVersionsCmd)
	getFileCmd.AddCommand(getFileRetentionPoliciesCmd)
	getFileCmd.AddCommand(getFileCollectionsCmd)

	createFileCmd.Flags().StringVar(&fileName, "name", "", "File name")
	createFileCmd.Flags().StringVar(&fileDescription, "description", "", "File description")
	createFileCmd.Flags().StringVar(&fileMetadata, "metadata", "", "File metadata")
	createFileCmd.Flags().StringVar(&fileStorageService, "storage-service", "", "Service used to store file")
	createFileCmd.Flags().StringVar(&filePathUpload, "path", "", "Local file path to upload (required)")
	createFileCmd.MarkFlagRequired("path")

	attachFileCmd.Flags().StringVar(&fileName, "name", "", "File name")
	attachFileCmd.Flags().StringVar(&fileStorageService, "storage-service", "", "Service used to store file")
	attachFileCmd.Flags().StringVar(&fileStorageId, "storage-id", "", "Id of file in existing storage service")

	copyFileCmd.Flags().StringVar(&fileStorageService, "storage-service", "", "Service used to store file")

	updateFileCmd.Flags().StringVar(&fileName, "name", "", "File name")
	updateFileCmd.Flags().StringVar(&fileDescription, "description", "", "File description")
	updateFileCmd.Flags().StringVar(&fileMetadata, "metadata", "", "File metadata")

	addVersionCmd.Flags().StringVar(&fileName, "name", "", "File name")
	addVersionCmd.Flags().StringVar(&fileStorageService, "storage-service", "", "Service used to store file")
	addVersionCmd.Flags().StringVar(&filePathUpload, "path", "", "Local file path to upload (required)")
	addVersionCmd.MarkFlagRequired("path")

	getDirectUploadCmd.Flags().StringVar(&fileStorageService, "storage-service", "", "Service used to store file")
	getDirectUploadCmd.Flags().Int32Var(&directUploadExpires, "expires-in", 0, "When upload link expires in seconds")

	downloadFileCmd.Flags().StringVar(&fileName, "name", "", "Change file name on local system when downloading")
	downloadFileCmd.Flags().StringVar(&filePathDownload, "path", "", "Local file path to download (required)")
	downloadFileCmd.MarkFlagRequired("path")
}
