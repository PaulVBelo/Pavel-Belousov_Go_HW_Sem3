package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"math/rand"
	"task2/dto"
)

type Server struct {
	*http.Server
}

func NewServer(addr string) *Server{
	srv := &http.Server{
		Addr: addr,
	}

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprint(w, "v1.0.0")
	})

	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		// Есть ли способ это проверять поприятнее? 
		// В плане в том же SpringBoot есть аннотации @GetMapping, @PostMapping и т.п. (Хотя там много всего упрощено :D)
		// Для го нет таких аналогов?
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req dto.DecodeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		b, err := base64.StdEncoding.DecodeString(req.InputString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		outputString := string(b)
		resp := dto.DecodeResponse{OutputString: outputString}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/hard-op", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		time.Sleep(time.Duration(rand.Intn(10)+10) * time.Second)

		status := rand.Intn(2)
		if status == 0 {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Oppeartion successful")
	})

	return &Server{srv}
}

func (s *Server) Start() error {
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
