package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Container struct {
	ID     string
	Name   string
	Status string
}

func ListContainer() ([]Container, error) {
	cmd := exec.Command("docker", "ps", "--format", "{{.ID}}|{{.Names}}|{{.Status}}")
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(string(output))
	if trimmed == "" {
		fmt.Println("Active container not found")
		return []Container{}, nil
	}
	lines := strings.Split(trimmed, "\n")

	var containers []Container
	for _, line := range lines {
		parts := strings.Split(line, "|")
		container := Container{
			ID:     parts[0],
			Name:   parts[1],
			Status: parts[2],
		}
		containers = append(containers, container)
	}

	return containers, nil
}
