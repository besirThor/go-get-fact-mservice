package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// cat fact service

// * Service is an interface that defines a method for retrieving cat facts.
type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

// * CatFactService is a type that implements the Service interface.
type CatFactService struct {
	url string // the url is for the cat fact api
}

// * NewCatFactService is a constructor function that returns a new instance of the CatFactService type.

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

// * GetCatFact retrieves a cat fact from the API specified by the url field of the CatFactService instance.
func (s *CatFactService) GetCatFact(ctx context.Context) (*CatFact, error) {
	resp, err := http.Get(s.url) // Send an HTTP GET request to the API.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fact := &CatFact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil { // * Decode the response body into a CatFact struct.
		return nil, err
	}
	return fact, nil //return the cat fact
}
