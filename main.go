package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

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

func main() {
	//create new router object
	r := mux.NewRouter()

	// http.HandleFunc("/", handler)
	fmt.Println("Registering / path handler")
	r.HandleFunc("/", homeHandler).Methods("GET")

	fmt.Println("Registering /users/ path handler")
	handlers.GetUser(r)

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", r)
}
