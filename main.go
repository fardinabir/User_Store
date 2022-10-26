package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User ...
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

// Articles ...
var Users []User

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", createNewUsers).Methods("POST")
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers).Methods("GET")

	//myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	//myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Users)
}

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnSingleArticle")

// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	for _, article := range Articles {
// 		if article.Id == key {
// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}
// }

func createNewUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewUsers")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(reqBody)
	var newUser User
	json.Unmarshal(reqBody, &newUser)
	Users = append(Users, newUser)

	fmt.Println(Users)
	json.NewEncoder(w).Encode(newUser)
}

// func deleteArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: deleteArticle")
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	for index, article := range Articles {
// 		if article.Id == id {
// 			Articles = append(Articles[:index], Articles[index+1:]...)
// 		}
// 	}

// }

func main() {
	Users = []User{
		User{
			FirstName: "fardin",
			LastName:  "abir",
			Password:  "1234",
			Phone:     "131424243",
		},
	}

	handleRequests()
}
