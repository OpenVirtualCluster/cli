package cmd

import (
	"github.com/OpenVirtualCluster/cli/assets"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"net/http"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Open the dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		serveDashboard()
	},
}

func serveDashboard() {
	fs, err := fs.Sub(assets.NextFS, ".next")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(fs)))
	log.Println("Starting server at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
