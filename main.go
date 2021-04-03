package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article

	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var article Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &article)
	for index, artcl := range Articles {
		if artcl.Id == id {
			Articles[index].Title = article.Title
			Articles[index].Desc = article.Desc
			Articles[index].Content = article.Content
		}
		break
	}
}

func handleRequests() {
	m := mux.NewRouter().StrictSlash(true)

	m.HandleFunc("/", homepage)
	m.HandleFunc("/articles", returnAllArticles)
	m.HandleFunc("/article", createNewArticle).Methods("POST")
	m.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	m.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	m.HandleFunc("/article/{id}", returnArticleById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", m))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello1", Desc: "Hello Article 1", Content: "Hello Article 1 Content"},
		Article{Id: "2", Title: "Hello2", Desc: "Hello Article 2", Content: "Hello Article 2 Content"},
	}
	handleRequests()
}
