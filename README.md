# Ascii Art Web

## Description
This project is a simple web application that generates ASCII art based on user input. Users can enter text and select a banner style, and the application will render the corresponding ASCII art.

In this version a Download button is added, which allows the users to download the generated ascii art as a txt file to their computer.

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
If the POST request comes from the download button (which only appears on the page when ascii art has been generated), the generated ascii art is saved as a txt file on the user's computer. 

**ASCII Art Generation:**

Upon valid input, the server calls the generateAsciiArt function.
This function takes the user input (text) and processes it to create ASCII art.

**Template Rendering:**

The server uses Go's html/template package to render the HTML page.
It dynamically injects the generated ASCII art into the HTML template.

**Response:**

After processing, the server sends the rendered page back to the user, displaying the ASCII art and download button.

## Changelog 
- MakeAsciiArtTable now checks that the banner file is the correct length. It returns an error if the banner file can't be read or the line number is incorrect.
- In addition to handling errors received from MakeAsciiArtTable, GenerateAsciiArt also catches any unexpected panic and recovers. In both cases, an error is sent to the MainPage function, which then triggers an Internal Server Error. Under normal circumstances there should never be a panic, but in order to test this there is a deliberate panic trigger in the code that can be commented out.
- Added export functionality by including a download button in the template.
