package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/random-customer-service/api"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", randomUserHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" // Default to 3001
		log.Printf("Defaulting to port %s", port)
	}
	fmt.Println("Server is running on http://localhost:", port)
	http.ListenAndServe(":"+port, nil)
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
