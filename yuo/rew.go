package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.Handle("/", http.FileServer(http.Dir("./templates")))

	direccion := ":8080"
	fmt.Println("escuchando..." + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
	
}