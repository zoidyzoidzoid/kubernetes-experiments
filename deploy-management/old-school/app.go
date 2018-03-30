package main

import (
	"fmt"
	"gopkg.in/redis.v5"
	"log"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Incr("hits").Err()
	if err != nil {
		log.Print(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	hits, err := client.Get("hits").Result()
	if err != nil {
		log.Print(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Hello World! I have been seen %s times.", hits)
}

func serverNameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, hostname)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/server-name", serverNameHandler)
	host := os.Getenv("DOCKER_COUNTER_ADDR")
	if host == "" {
		host = "127.0.0.1"
	}
	fmt.Printf("Now starting server on http://%s:8080\n", host)
	http.ListenAndServe(host+":8080", nil)
}
