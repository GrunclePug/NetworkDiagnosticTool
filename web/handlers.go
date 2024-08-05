package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// User Holds form data
type User struct {
	Username string
}

// UserHandler Handles HTTP Requests for User form
func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] New Request, Target: %v Source: %v\n", r.URL.Path, r.RemoteAddr)

	switch r.URL.Path {
	case "/": // Index
		fmt.Println("[INFO] Serving Index")
		templating(w, "web/static/index.html", nil)
	case "/send": // Form Submission
		if r.Method == "POST" {
			// Parse Form
			fmt.Println("[INFO] POST Received to /send")
			err := r.ParseForm()
			if err != nil {
				fmt.Println("[ERROR] Parse Form Error: ", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else { // GET
			fmt.Println("[INFO] GET Received to /send")
		}

		// Pull Form Value
		username := r.FormValue("user")
		fmt.Println("[INFO] Parsed User from Request: ", username)

		// Build HTML Template for response
		user := User{
			Username: username,
		}
		fmt.Println("[INFO] Sending Response: ", user)
		templating(w, "web/static/index.html", user)
	}
}

// templating HTML Serving using Templates
func templating(w http.ResponseWriter, fileName string, data interface{}) {
	// Read HTML file
	t := template.Must(template.ParseFiles(fileName))

	// Serve HTML data
	err := t.Execute(w, data)
	if err != nil {
		fmt.Println("[ERROR] HTML Template Execution Error: ", err.Error())
		log.Fatal(err)
	}
}
