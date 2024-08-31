package api

import (
	"log"
	"net/http"
	"zeorgo/middleware"
)

type APIServer struct {
	addr string
}

func NewAPIServer(address string) *APIServer {
	return &APIServer{
		addr: address,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID" + userID))
	})
	// v1 := http.NewServeMux()
	// v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))
	// multi middleware
	middlewareChain := middleware.CollectMiddlewares()
	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}
	log.Print("Server is running")
	return server.ListenAndServe()
}
