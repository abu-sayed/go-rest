package main

import (
	"net/http"
)

func main()  {
	http.ListenAndServe(":9000", GetRouter()) // Server started listening at localhost:9000
}
