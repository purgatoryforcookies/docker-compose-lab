package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Endpoint struct {
	url     string
	healthy bool
}

var url = Endpoint{url: "/health", healthy: true}

func health(w http.ResponseWriter, r *http.Request) {
	if url.healthy {
		fmt.Fprintf(w, "hello\n")
	} else {
		http.Error(w, "Not healthy at all", http.StatusBadRequest)
	}
}

func (e *Endpoint) changeEndpoint(s int) {
	seconds := time.Duration(s)

	time.Sleep(seconds * time.Second)
	fmt.Println("Container 1 becoming unhealthy")
	e.healthy = false
}

func main() {

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	input, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		panic("Cannot convert to integer")
	}
	seconds := time.Duration(input)

	fmt.Printf("Waiting for %s \n", seconds*time.Second)
	time.Sleep(seconds * time.Second)
	go url.changeEndpoint(10)

	fmt.Printf("Running a web server in %s \n", port)

	http.HandleFunc(url.url, health)
	http.ListenAndServe(port, nil)

}
