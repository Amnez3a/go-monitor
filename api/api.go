package api

import (
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
	// only post-method [TEST]
	if r.Method != http.MethodPost {
		fmt.Println(w, "only post-method", http.StatusMethodNotAllowed)
		return
	}
	var statusResult []StatusServer
	err := json.NewDecoder(r.Body).Decode(&statusResult)
	if err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8") // default text
	for _, server := range statusResult {
		statusText := "offline"
		if server.Online {
			statusText = "online"
		}
		fmt.Fprintf(w, "%s [%s] : [%s]\n", server.Name, server.Address, statusText)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Status: %d", http.StatusOK)
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
