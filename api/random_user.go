package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetRandomUser() string {
	// Define the URL
	url := "https://randomuser.me/api/"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	// Return response body
	return string(body)

}
