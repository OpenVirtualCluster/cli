package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Version variable, to be set at build time
var Version = "v0.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OVC",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("OVC CLI version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
