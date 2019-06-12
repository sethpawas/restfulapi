package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type articles struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type e1 []articles

var (
	mutex sync.Mutex
	s1    e1
)

func update(s4 e1, s3 articles, a int, w http.ResponseWriter) {
	mutex.Lock()
	defer mutex.Unlock()
	for index, articles := range s1 {
		if articles.Age == a {
			s4 = append(s1[:index], s3)
			//json.NewEncoder(w).Encode(s4v)
			s4 = append(s4, s1[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(s4)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handlerequests() {
	myrouter := mux.NewRouter().StrictSlash(true)

	myrouter.HandleFunc("/", homepage)
	myrouter.HandleFunc("/fetch", fetcharticles)
	myrouter.HandleFunc("/add", addarticles).Methods("POST")
	myrouter.HandleFunc("/search/{age}", returnsinglearticles)
	myrouter.HandleFunc("/delete/{age}", deletearticles)
	myrouter.HandleFunc("/update/{age}", updatearticles).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myrouter))

}
func fetcharticles(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(s1)
	fmt.Println("endpoint hi Homepage")

}
func returnsinglearticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["age"]
	a, _ := strconv.Atoi(key)

	for _, articles := range s1 {
		if articles.Age == a {
			json.NewEncoder(w).Encode(articles)
		}

	}

	//fmt.Fprintf(w, "Key: "+key)
}
func addarticles(w http.ResponseWriter, r *http.Request) {
	Str, _ := ioutil.ReadAll(r.Body)
	var s2 articles
	json.Unmarshal(Str, &s2)
	s1 = append(s1, s2)
	fmt.Println(s1)
	json.NewEncoder(w).Encode(s2)

}
func deletearticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["age"]
	a, _ := strconv.Atoi(key)
	fmt.Printf("delete hit")
	for index, articles := range s1 {

		if articles.Age == a {

			s1 = append(s1[:index], s1[index+1:]...)
			if index == 0 {
				fmt.Fprint(w, "last element deletion query")
			}
		}
	}
}
func updatearticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["age"]
	a, _ := strconv.Atoi(key)
	str, _ := ioutil.ReadAll(r.Body)
	var s3 articles
	var s4 e1
	json.Unmarshal(str, &s3)
	update(s4, s3, a, w)
}
func main() {
	s1 = e1{
		articles{Name: "ayushi", Age: 20},
		articles{Name: "ayush", Age: 10},
	}
	handlerequests()
}
