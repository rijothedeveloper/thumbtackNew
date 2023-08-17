package categorycontroller

import (
	"io"
	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world! from category\n")
}