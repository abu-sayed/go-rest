package main

import (
	"net/http"
	"github.com/abu-sayed/go-rest/services/db"
)

func main() {
	// Server started listening at localhost:9000
	http.ListenAndServe(":9000", GetRouter())
	defer db.GetDb().Close()
}
