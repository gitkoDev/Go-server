package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		fmt.Println("sa")
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello")
}

func formPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "POST request successful")
	}
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "hello")
	fmt.Fprintln(w, name, address, "data here", r.Form)

}

func main() {
	srv := http.FileServer(http.Dir("./static"))

	http.Handle("/", srv)
	http.HandleFunc("/hello", helloPage)
	http.HandleFunc("/form", formPage)

	fmt.Println("running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
