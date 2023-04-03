package main

import (
	"log"
)

func main() {

	// *create a new CatFactService instance that fetches cat facts from https://catfact.ninja/fact
	svc := NewCatFactService("https://catfact.ninja/fact")
	// *wrap the CatFactService instance with a LoggingService instance to log requests and responses
	svc = NewLoggingService(svc)

	//* create a new ApiServer instance with the wrapped service
	apiServer := newApiServer(svc)

	//* start the server and listen on port 8080
	log.Fatal(apiServer.Start(":8080"))
}
