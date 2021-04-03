package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page!!!")
	fmt.Println("Endpoint hit: Homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	m := mux.NewRouter().StrictSlash(true)

	m.HandleFunc("/", homepage)
	m.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8082", m))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello1", Desc: "Hello Article 1", Content: "Hello Article 1 Content"},
		Article{Title: "Hello2", Desc: "Hello Article 2", Content: "Hello Article 2 Content"},
	}
	handleRequests()
	//	fmt.Println("Hello world")
}
