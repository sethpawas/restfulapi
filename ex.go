package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
) //mux allow us easily help us to retrieve path and query parameters

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//used to get populated later in main function from database
type Articles []Article

var a1 Articles

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return all articles")
	json.NewEncoder(w).Encode(a1) //does the encoding our articles array into a JSON string and then writing as partr of our response
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range a1 {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	//get the body of our post request
	//return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	//unmarshal to the new article struct
	var a2 Article
	json.Unmarshal(reqBody, &a2)
	//update the global array
	a1 = append(a1, a2)

	json.NewEncoder(w).Encode(a2)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	//parse the path parameters
	vars := mux.Vars(r)
	//extract the id need to delete
	id := vars["id"]

	for index, article := range a1 {
		if article.Id == id {
			a1 = append(a1[:index], a1[index+1:]...)
		}
	}
}

//handle all requests to our root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

//match the Url path hit with a defined function
func handleRequests() {

	//ordering is important
	//new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//replace http. with myRouter
	myRouter.HandleFunc("/", homePage)
	//map any call to /articles
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	//we will pass the newly instance instead of nil
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

//kickoff our API
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	a1 = Articles{
		Article{Id: "1", Title: "Hello", Desc: "Description details", Content: "Content Description"},
		Article{Id: "2", Title: "Hello 2", Desc: "Description details", Content: "Content Description"},
	}
	handleRequests()
}
