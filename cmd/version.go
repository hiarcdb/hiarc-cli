package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hiarc CLI",
	Long:  `All software has versions. This is Hiarc's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hiarc CLI v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
