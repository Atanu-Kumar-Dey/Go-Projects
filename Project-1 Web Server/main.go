package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Simple Static Web Server")
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Bad request <h1>404 not found</h1>", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not compatable ", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello and greetings to all golang user")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Error while submitting the form")
		return
	}
	fmt.Fprintf(w, "POST request successful")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	fmt.Fprintf(w, "firstname = %s\n", firstname)
	fmt.Fprintf(w, "lastname = %s\n", lastname)
}
