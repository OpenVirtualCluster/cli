package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/OpenVirtualCluster/cli/assets"
	"github.com/spf13/cobra"
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
	var helmData []byte
	var helmFilename string

	switch runtime.GOOS {
	case "darwin":
		helmData = assets.HelmDarwin
		helmFilename = "helm-darwin"
	case "linux":
		helmData = assets.HelmLinux
		helmFilename = "helm-linux"
	default:
		log.Fatalf("Unsupported OS: %v", runtime.GOOS)
	}

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
