package cmd

import (
	"go-monitor-docker/internal/monitor"
	"time"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		interval, _ := cmd.Flags().GetDuration("interval")
		monitor.Watch(interval)
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.Flags().DurationP("interval", "i", 5*time.Second, "reload containers status")
}
