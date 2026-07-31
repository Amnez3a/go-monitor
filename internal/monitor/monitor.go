package monitor

import (
	"context"
	"fmt"
	"go-monitor-docker/internal/docker"
	"os"
	"os/signal"
	"time"
)

func clearScreen() { fmt.Print("\033[H\033[2J") }

func Watch(interval time.Duration) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("shutting down..")
			return

		case <-ticker.C:
			containers, err := docker.ListContainer()
			if err != nil {
				fmt.Println("E: ", err)
				continue
			}
			clearScreen()
			fmt.Println("Watch mode\n----------------")
			for _, c := range containers {
				fmt.Printf("| %s | %s | %s |\n", c.ID, c.Name, c.Status)
			}
		}
	}
}
