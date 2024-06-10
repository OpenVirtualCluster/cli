package cmd

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/OpenVirtualCluster/cli/assets"
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Open the dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		serveDashboard()
	},
}

func serveDashboard() {
	tempDir, err := os.MkdirTemp("", "ovc-dashboard")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := extractDashboard(tempDir); err != nil {
		log.Fatalf("Failed to extract dashboard: %v", err)
	}

	http.Handle("/", http.FileServer(http.Dir(tempDir)))
	log.Println("Serving dashboard at http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func extractDashboard(destination string) error {
	zipReader, err := zip.NewReader(bytes.NewReader(assets.DashboardZip), int64(len(assets.DashboardZip)))
	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		filePath := filepath.Join(destination, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
