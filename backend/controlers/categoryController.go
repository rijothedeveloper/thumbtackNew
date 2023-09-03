package controlers

import (
	"encoding/json"
	"net/http"
	"thumtack_category/services"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// io.WriteString(w, "Hello, world! from category\n")
	categories := services.GetCategories();
	// json, err := json.Marshal(&categories)
	// if err != nil {
	// 	io.WriteString(w, "error in converting categories to json\n")
	// }
	// io.WriteString(w, string(json))
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}