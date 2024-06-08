# Ascii Art Web

Transform your text into beautiful ASCII art with customizable options using Web Interface.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Ascii Art Algorith](#algorithm-description-for-ascii-art-generation)
- [Handling Errors](#handling-errors)
- [Contributing](#contributing)
- [Authors](#authors)

## Overview

Ascii Art Web is a webpage written in [Go] , [HTML] and [CSS]. It allows you to convert plain text into ASCII art, with the option to customize the text by choosing the banner you want.

## Installation

To install Ascii Art Web, you can use `go get`:

```bash
go get github.com/0xZED88/ascii-art-Web
```

## Usage

To use Ascii Art Web, run the following command:

```bash
go run main.go
```
Then you will start the server, check the terminal you will see something like that:
```bash
Server is Working succesfully....
```

Go to your Browser : 
- [localhost:8080](http://localhost:8080)

## Algorithm Description for ASCII Art Generation

1. **Input Handling:**
   - User submits text and selects one of three font styles via an HTML form.
   - The server parses the form data from an HTTP POST request.

2. **Validation:**
   - The server checks for missing text or font style inputs.
   - Validates the selected font style.

3. **ASCII Art Generation:**
   - Loads ASCII patterns for the selected banner.
   - Converts each character of the input text to its corresponding ASCII pattern.
   - Assembles these patterns line by line to form the complete ASCII art.

4. **Template Rendering:**
   - Loads an HTML template for displaying the ASCII art.
   - Inserts the generated ASCII art into the template.
   - Sends the final HTML response to the user’s browser.


## Handling Errors
This web App handled so far errors : 
- 400 in case of bad request
- 500 in case the server has some issues in respnod
- 404 if the page or banner is not found
- 405 if Method not allowed 

## Authors
This Web App it's been programmed and Designed by : 
- Zayd Zitan
- Othman Qritel
- Mohamed Tawil

## Contributing

Contributions are welcome! Please feel free to submit bug reports, feature requests, or pull requests.
