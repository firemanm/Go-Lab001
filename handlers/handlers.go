package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/firemanm/LAB001/errors"
	"github.com/firemanm/LAB001/models"
	"github.com/gorilla/mux"
)

var users = make(map[string]models.User)

func GetUser(router *mux.Router) {
	router.HandleFunc("/users/{id}", loggingMiddleware(getUser)).Methods("GET")
	fmt.Println("/users/ path registered")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user, ok := users[id]
	if !ok {
		err := errors.NewAppError("User not found", http.StatusNotFound)
		http.Error(w, err.Error(), err.Code)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		elapsed := time.Since(start)
		log.Printf("%s %s %v\n", r.Method, r.URL, elapsed)
	}

}
