package server

type Checker interface {
	Check(s Server) string
}

type Server struct {
	Name    string `json:"name"`
	IP      string `json:"ip"`
	Port    int    `json:"port"`
	Status  string
	Latency int
}
