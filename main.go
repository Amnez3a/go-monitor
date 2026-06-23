package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"go-monitor/checker"
	"go-monitor/server"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func runAllChecks(servers []server.Server, c checker.TCPChecker) {
	clearScreen()

	var wg sync.WaitGroup
	for _, srv := range servers {
		wg.Add(1)
		go func(s server.Server) {
			defer wg.Done()
			result := c.Check(s)
			fmt.Println(result)
		}(srv)
	}
	wg.Wait()

	fmt.Println("\n--------------------------------------------------")
	fmt.Println("Enter - reload stats")
	fmt.Println("Ctrl+\\ - exit")
}

func main() {
	file := flag.String("file", "servers.json", "path to file")
	timeout := flag.Int("timeout", 3, "timeout of seconds")
	flag.Parse()

	data, err := os.ReadFile(*file)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	var servers []server.Server
	err = json.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := checker.TCPChecker{Timeout: time.Duration(*timeout) * time.Second}
	runAllChecks(servers, c)
	for {
		fmt.Scanln()
		runAllChecks(servers, c)
	}
}
