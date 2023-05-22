package main

import (
	"fmt"
	types "github.com/franciscocunha55/apiGO/types"
	"github.com/gorilla/mux"
	"net/http"
)

var users []types.User

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler)
	router.HandleFunc("/{id}", getUser).Methods("GET")

	http.Handle("/", router)

	userManel := types.NewUser("manel", 30)
	fmt.Println(userManel.Name)
	types.ChangeName(userManel, "Ramiro")
	fmt.Println(userManel.Name)

	types.
		fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", router)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Worldddd!")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "Getting user with ID: %s\n", id)
}
