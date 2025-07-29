package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func asyncTask(taskID int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Printf("Task %d completed\n", taskID)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go asyncTask(i, &wg)
	}
	wg.Wait()
	fmt.Fprintln(w, "All tasks completed")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
