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

type employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type e1 []employee

var (
	mutex sync.Mutex
	s1    e1
)

func update(s4 e1, s3 employee, a int, w http.ResponseWriter) {
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
	myrouter.HandleFunc("/fetch", fetchemployee)
	myrouter.HandleFunc("/add", addemployee).Methods("POST")
	myrouter.HandleFunc("/search/{age}", returnsingleemployee)
	myrouter.HandleFunc("/delete/{age}", deleteemployee)
	myrouter.HandleFunc("/update/{age}", updateemployee).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", myrouter))

}
func fetchemployee(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(s1)
	fmt.Println("endpoint hi Homepage")

}
func returnsingleemployee(w http.ResponseWriter, r *http.Request) {
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
func addemployee(w http.ResponseWriter, r *http.Request) {
	Str, _ := ioutil.ReadAll(r.Body)
	var s2 employee
	json.Unmarshal(Str, &s2)
	s1 = append(s1, s2)
	fmt.Println(s1)
	json.NewEncoder(w).Encode(s2)

}
func deleteemployee(w http.ResponseWriter, r *http.Request) {
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
func updateemployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["age"]
	a, _ := strconv.Atoi(key)
	str, _ := ioutil.ReadAll(r.Body)
	var s3 employee
	var s4 e1
	json.Unmarshal(str, &s3)
	update(s4, s3, a, w)
}
func main() {
	s1 = e1{
		employee{Name: "ayushi", Age: 20},
		employee{Name: "ayush", Age: 10},
	}
	handlerequests()
}
