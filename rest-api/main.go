package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article 

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title:"art 1", Desc: "Desc 1", Content: "cont 1"},
		Article{Title:"art 2", Desc: "Desc 2", Content: "cont 2"},
	}
	fmt.Println("Articles Endopoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}