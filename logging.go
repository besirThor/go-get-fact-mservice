package main

import (
	"context"
	"fmt"
	"time"
)

// * LoggingService is a type that wraps a Service and logs information about each request.
type LoggingService struct {
	next Service
}

// * NewLoggingService is a constructor function that returns a new instance of the LoggingService type.
func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

// * GetCatFact retrieves a cat fact from the wrapped Service and logs information about the request.
func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) { // Use a defer statement to log information about the request after it completes.
		fmt.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx) // *Retrieve the cat fact from the wrapped Service.
}
