package main

import (
	"fmt"
	"net/http"

	"github.com/random-customer-service/api"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", randomUserHandler)
	fmt.Println("Server is running on http://localhost:3001")
	http.ListenAndServe(":3001", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Welcome to Random user")
}
func randomUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		fmt.Fprintln(w, api.GetRandomUser())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
