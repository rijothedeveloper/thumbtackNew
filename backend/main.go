package main

import (
	"log"
	"net/http"

	"thumtack_category/config"
	controler "thumtack_category/controlers"
)

func init() {
	config.ConnectDB();
}


func main() {
	http.HandleFunc("/categories", controler.GetCategories)
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Fatal("error starting server")
		return
	} 
}

// func getHome(w http.ResponseWriter, r *http.Request){
// 	io.WriteString(w, "Hello, world!\n")
// }
