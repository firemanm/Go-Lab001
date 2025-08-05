package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"encoding/json"
	"log"

	"github.com/firemanm/LAB001/handlers"
	"github.com/gorilla/mux"
)

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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
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
	//create new router object
	r := mux.NewRouter()

	// http.HandleFunc("/", handler)
	fmt.Println("Registering / path handler")
	r.HandleFunc("/", homeHandler).Methods("GET")

	fmt.Println("Registering /users/ path handler")
	handlers.GetUser(r)

	fmt.Println("Registering /health path handler")
	r.HandleFunc("/health", loggingMiddleware(healthHandler)).Methods("GET")


	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", r)
}
