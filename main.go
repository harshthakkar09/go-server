package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler for form
func formHandler(w http.ResponseWriter, r *http.Request) {
	// parsing the http form so that we can get values
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	// extracting value of name and address
	name := r.FormValue("name")
	address := r.FormValue("address")
	// printing values of name and address
	fmt.Fprintf(w, "Name= %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)
}

// handler to print hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// checking whether path is proper or not (can ignore as it would be call if and only if path is "/hello")
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// checking whether request method is GET or not
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	// printing hello
	fmt.Fprint(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)            // defined default path handler
	http.HandleFunc("/form", formHandler)   // defined hadler for form
	http.HandleFunc("/hello", helloHandler) // defined handler to print hello

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
