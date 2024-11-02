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

func main() {
	http.HandleFunc("/", mainPage)

	// Serve static files from the "static" folder
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	vars := PageVariables{}

	if r.Method == http.MethodGet {
		if r.URL.Path != "/" && r.URL.Path != "/index.html" {
			// Serve a custom 404 error page
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "templates/404.html")
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

		//handle download button
		if r.FormValue("download") == "true" {
			// User requested a download
			content := []byte(r.FormValue("ascii_art"))
			log.Println(vars.AsciiArt)

			// Set headers for file download
			w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
			w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
			w.Header().Set("Content-Type", "text/plain")

			_, err := w.Write(content)
			if err != nil {
				http.Error(w, "Failed to export ASCII art", http.StatusInternalServerError)
				log.Printf("Error writing ASCII art to response: %v", err)
			}
			return
		}

		vars.Text = r.FormValue("text")
		vars.Banner = r.FormValue("banner")

		// Validate text
		if vars.Text == "" {
			http.Error(w, "Bad Request 400: 'text' cannot be empty", http.StatusBadRequest)
			return
		}

		//Comment this out to test internal server error handling
		// Validate banner
		if vars.Banner != "standard" && vars.Banner != "shadow" && vars.Banner != "thinkertoy" {
			http.Error(w, "Bad Request 400: Invalid 'banner' value", http.StatusBadRequest)

			return
		}

		// Log received values for debugging
		log.Printf("Received text: %q", vars.Text)
		log.Printf("Selected banner: %q", vars.Banner)

		// Generate ASCII art from the input text
		var err error
		vars.AsciiArt, err = ascii.GenerateAsciiArt(vars.Text, vars.Banner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
