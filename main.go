package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	// renaming our package
	myFunc "ascii/ascii"
)

type Status struct {
	Code   string
	Statu  string
	Result string
}

var (
	status  Status
	ResData Status // re
)

/*type result struct {
	Result string
}
var re result
*/

func StartServer() {
	// Creating Local Server
	fmt.Println("Server is Working succesfully....")
	fmt.Println()
}

func index(w http.ResponseWriter, r *http.Request) {
	// Check the reciving method is it valid or not
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Handling the wrong path error.
	if r.URL.Path != "/" {
		status.Code = "404"
		status.Statu = "Page Not Found"
		// Changing status of the Header

		w.WriteHeader(http.StatusNotFound)
		temp, err := template.ParseFiles("./template/404.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, status)
		return
	}
	// make the template and parse it
	temp, err := template.ParseFiles("./template/index.html")
	// handlling the server error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func asciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	// Getting Form Values Text And Banner :)
	text := r.Form.Get("name")
	banner := r.Form.Get("banner")
	if text == "" || banner == "" {
		status.Code = "400"
		status.Statu = "Bad Request"
		w.WriteHeader(http.StatusBadRequest)
		temp, _ := template.ParseFiles("./template/404.html")
		temp.Execute(w, status)
		return
	}
	// Proccessing Data and Getting ascii art result
	Result, er := myFunc.Batata(banner, text)
	ResData.Result = Result
	// Handling the case if the banner is not exists
	if er {
		status.Code = "404"
		status.Statu = "Banner Not Found"
		w.WriteHeader(http.StatusNotFound)
		temp, _ := template.ParseFiles("./template/404.html")
		temp.Execute(w, status)
		return
	}

	// Creating new Object from the template
	temp, errr := template.ParseFiles("./template/ascii-art.html")
	if errr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, template.HTML(Result))
}

func downlaod(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
	w.Header().Set("Content-Type", "text/html")
	// set the lenght
	w.Header().Set("Content-Length", strconv.Itoa(len(status.Result)))
	// Serve the file
	ResData.Result = strings.ReplaceAll(ResData.Result, "<br>", "\n")
	// http.ServeFile(w, r, re.Result)
	w.Write([]byte(ResData.Result))
}

func main() {
	// Handling Home (root) Page
	http.HandleFunc("/", index)
	// Handling /Ascci-art path
	http.HandleFunc("/ascii-art", asciiArt)
	http.HandleFunc("/downloads/", downlaod)
	// := strings.Replace("\n", "<br>")
	// host style file to serve it later
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	// we intialized the local server with 8080 port
	http.ListenAndServe(":8080", nil)
}
