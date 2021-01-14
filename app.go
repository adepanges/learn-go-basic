package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello 1", Desc: "Article Description 1", Content: "Article Content 1"},
		Article{Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
	fmt.Fprintf(w, "articles")
	json.NewEncoder(w).Encode(articles)
	fmt.Println("Endpoint Hit: homePage", articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getArticles)
	log.Fatal(http.ListenAndServe(":8090", nil))
	fmt.Println("ListenAndServe 8090")
}

func main() {
	handleRequests()
}
