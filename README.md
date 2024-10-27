# Ascii Art Web

## Description
This project is a simple web application that generates ASCII art based on user input. Users can enter text and select a banner style, and the application will render the corresponding ASCII art.
The servers runs at localhost:8080

## Authors
- George Choremis
- Arsen Tsntsgoukian
- Katerina Vamvasaki

## Implementation Details

The application consists of a simple flow to handle user input and generate ASCII art:

**Input Handling:**

The server listens for POST requests containing user input (text and banner selection).
When a POST request is received, the server parses the form data.

**ASCII Art Generation:**

Upon valid input, the server calls the generateAsciiArt function.
This function takes the user input (text) and processes it to create ASCII art.

**Template Rendering:**

The server uses Go's html/template package to render the HTML page.
It dynamically injects the generated ASCII art into the HTML template.

**Response:**

After processing, the server sends the rendered page back to the user, displaying the ASCII art.