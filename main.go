package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-monitor/api"
	"go-monitor/checker"
	"go-monitor/server"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

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
	// web
	a := &api.API{
		Servers: servers,
		Checker: c,
	}
	go api.StartServer(a)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
