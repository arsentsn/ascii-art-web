package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/ascii"
)

// PageVariables holds data to pass to the HTML template
type PageVariables struct {
	Text     string
	Banner   string
	AsciiArt string
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	vars := PageVariables{}

	if r.Method == http.MethodGet {
		if r.URL.Path != "/" && r.URL.Path != "/index.html" {
			handler404(w, r)
			return
		}
	}

	if r.Method == http.MethodPost {
		// Handle POST data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing form: %v", err)
			return
		}

		vars.Text = r.FormValue("text")
		vars.Banner = r.FormValue("banner")

		// Validate text
		if vars.Text == "" {
			http.Error(w, "Bad Request 400: 'text' cannot be empty", http.StatusBadRequest)
			return
		}

		// Validate banner
		if vars.Banner != "standard" && vars.Banner != "shadow" && vars.Banner != "thinkertoy" {
			http.Error(w, "Bad Request 400: Invalid 'banner' value", http.StatusBadRequest)

			return
		}

		// Log received values for debugging
		log.Printf("Received text: %q", vars.Text)
		log.Printf("Selected banner: %q", vars.Banner)

		// Generate ASCII art from the input text
		vars.AsciiArt = ascii.GenerateAsciiArt(vars.Text, vars.Banner)

	}
	// Serve the main template for the homepage
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}
	// Set status code to OK (200)
	w.WriteHeader(http.StatusOK)
	if err := t.Execute(w, vars); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)

	}
}

func handler404(w http.ResponseWriter, r *http.Request) {
	// Serve a custom 404 error page
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "templates/404.html")
}

func main() {
	http.HandleFunc("/", mainPage)

	http.HandleFunc("/404", handler404)

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
