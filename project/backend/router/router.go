package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type channel struct {
	Name string   `json:name`
	User username `json:user`
}
type username struct {
	ID   int    `json:ID`
	Name string `json:name`
}

type allUsername []username // list of usernames

var usernames = allUsername{
	{
		ID:   123,
		Name: "@bot",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

//func to view all usernames
func viewUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usernames)
}

//func to create an username
func createUsername(w http.ResponseWriter, r *http.Request) {
	var newUsername username
	reqBody, err := ioutil.ReadAll(r.Body) // read from the body

	//check for errors
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &newUsername)

	newUsername.ID = int(time.Now().Unix())   //cretate an unique id
	newUsername.Name = "@" + newUsername.Name // add an @ before every username

	usernames = append(usernames, newUsername) // add to the list of all username

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUsername)

}

//ConfigureRouter setup the router
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/viewUsername", viewUsername).Methods("GET")
	router.HandleFunc("/createUsername", createUsername).Methods("POST")

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
