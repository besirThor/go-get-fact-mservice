package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// * Define the ApiServer struct that holds an instance of the Service interface.
type ApiServer struct {
	svc Service
}

// * Define a constructor function for ApiServer that accepts an instance of the Service interface.
func newApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

// * Define the Start function for ApiServer that starts an HTTP server and registers the handleGetCatFact function.

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

// * Define the handleGetCatFact function that handles GET requests for a cat fact by invoking the GetCatFact method on the Service interface.

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		// If there is an error, return a JSON response with the error message and a status code of 422.
		writeJson(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}
	// Otherwise, return a JSON response with the cat fact and a status code of 200.

	writeJson(w, http.StatusOK, fact)
}

// *Define the writeJson function that writes a JSON response to the http.ResponseWriter with the specified status code and data.

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
