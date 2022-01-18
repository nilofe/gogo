package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	
)

type task struct {
	ID int "json:ID"
	Name string "json:Name"
	Content string "json:Content"
}
type allTasks []task 

var tasks = allTasks {
	{
		ID: 1,
		Name: "Task One",
		Content: "Some Content",
	},
}
func indexRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome dogs to my A-pi")
}


func main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/" , indexRoute)

	log.Println("escuchando...")
	http.ListenAndServe(":3000", router)
}
