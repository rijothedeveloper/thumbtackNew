package main

import (
	"log"
	"net/http"
	categorycontroller "thumtack_category/categoryController"
)

func main() {
	http.HandleFunc("/categories", categorycontroller.GetCategories)
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Fatal("error starting server")
		return
	} 
}

// func getHome(w http.ResponseWriter, r *http.Request){
// 	io.WriteString(w, "Hello, world!\n")
// }
