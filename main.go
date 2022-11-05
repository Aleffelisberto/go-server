package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(response, "Request sent successfully\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(response, "%s - %s", name, address)
}

func pageHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/page" {
		http.Error(response, "404 Not Found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(response, "Be welcome :)")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/page", pageHandler)
	fmt.Printf("Running server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}