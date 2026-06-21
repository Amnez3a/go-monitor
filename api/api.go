package api

import (
	"strings"
	"encoding/json"
	"fmt"
	"go-monitor/checker"
	"go-monitor/server"
	"net/http"
)

type StatusServer struct {
	Name    string `json:"name"`
	Online  bool   `json:"online"`
	Address string `json:"address"`
}

type Responce struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type API struct {
	Servers []server.Server
	Checker checker.TCPChecker
}

func (a *API) statusHandler(w http.ResponseWriter, r *http.Request) {
	var statusResult []StatusServer
	for _, s := range a.Servers {
		result := a.Checker.Check(s)
		fmt.Println("DEBUG:", result)
		online := strings.HasPrefix(result, "[✓]")
		statusResult = append(statusResult, StatusServer{
			Name:    s.Name,
			Online:  online,
			Address: s.IP,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"servers": statusResult,
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go!")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := Responce{Status: "ok", Message: "is working"}
	json.NewEncoder(w).Encode(resp)
}

func StartServer(a *API) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/status", a.statusHandler)
	http.HandleFunc("/health", health)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
