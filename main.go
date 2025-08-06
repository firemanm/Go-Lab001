package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/firemanm/LAB001/handlers"
	"github.com/gorilla/mux"
)

// create new router object
var router = mux.NewRouter()
var healthState = "OK"

func asyncTask(taskID int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Printf("Task %d completed\n", taskID)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go asyncTask(i, &wg)
	}
	wg.Wait()
	fmt.Fprintln(w, "All tasks completed")
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	healthState = "KillThemAll"
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if healthState == "OK" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(healthState)
	}

}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		elapsed := time.Since(start)
		log.Printf("%s %s %v\n", r.Method, r.URL, elapsed)
		fmt.Println(r.Method, r.URL, elapsed)

	}

}

func main() {
	// http.HandleFunc("/", handler)
	fmt.Println("Registering / path handler")
	router.HandleFunc("/", homeHandler).Methods("GET")

	fmt.Println("Registering /users/ path handler")
	handlers.GetUser(router)

	fmt.Println("Registering /health path handler")
	router.HandleFunc("/health", loggingMiddleware(healthHandler)).Methods("GET")

	fmt.Println("Registering /stop path handler")
	router.HandleFunc("/stop", loggingMiddleware(stopHandler)).Methods("GET")

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", router)
}
