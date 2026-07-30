package cmd

import (
	"fmt"

	"go-monitor-docker/internal/docker"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check containers status",
	Run: func(cmd *cobra.Command, args []string) {
		containers, err := docker.ListContainer()
		if err != nil {
			fmt.Println("E: ", err)
			return
		}
		for _, c := range containers {
			fmt.Printf("ID: %s | Name: %s | Status: %s\n", c.ID, c.Name, c.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
