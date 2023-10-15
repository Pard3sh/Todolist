package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color: steelblue'> Hello </h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := Book{
		// this is where we would query our database instead of hard coding
		Title:  "Giovanni's Room",
		Author: "James Baldwin",
		Pages:  176,
	}
	//makes an encoder with our response writer and then encode our book into that json
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)

	err := http.ListenAndServe(":5100", nil)
	if err != nil {
		log.Fatal(err)
	}
}
