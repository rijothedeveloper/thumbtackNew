package controlers

import (
	"encoding/json"
	"io"
	"net/http"
	"thumtack_category/services"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world! from category\n")
	categories := services.GetCategories();
	// json, err := json.Marshal(&categories)
	// if err != nil {
	// 	io.WriteString(w, "error in converting categories to json\n")
	// }
	// io.WriteString(w, string(json))
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(categories)
}