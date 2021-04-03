package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
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

func returnArticleById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
			break
		}
	}
}

func handleRequests() {
	m := mux.NewRouter().StrictSlash(true)

	m.HandleFunc("/", homepage)
	m.HandleFunc("/articles", returnAllArticles)
	m.HandleFunc("/article/{id}", returnArticleById)
	log.Fatal(http.ListenAndServe(":8082", m))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello1", Desc: "Hello Article 1", Content: "Hello Article 1 Content"},
		Article{Id: "2", Title: "Hello2", Desc: "Hello Article 2", Content: "Hello Article 2 Content"},
	}
	handleRequests()
}
