package cmd

import (
	"github.com/OpenVirtualCluster/cli/assets"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the virtual-cluster-operator Helm chart",
	Run: func(cmd *cobra.Command, args []string) {
		helmPath := extractHelmBinary()
		commands := [][]string{
			{helmPath, "repo", "add", "ovc-vco", "https://charts.openvirtualcluster.dev"},
			{helmPath, "repo", "update"},
			{helmPath, "install", "vco", "ovc-vco/virtual-cluster-operator"},
		}

		for _, command := range commands {
			cmd := exec.Command(command[0], command[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("Command failed: %v", err)
			}
		}
	},
}

func extractHelmBinary() string {
	var helmData []byte = assets.Helm
	var helmFilename string = "helm"

	tempDir, err := ioutil.TempDir("", "ovc-helm")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}

	helmPath := filepath.Join(tempDir, helmFilename)
	if err := ioutil.WriteFile(helmPath, helmData, 0755); err != nil {
		log.Fatalf("Failed to write helm binary: %v", err)
	}

	return helmPath
}

func init() {
	rootCmd.AddCommand(installCmd)
}
