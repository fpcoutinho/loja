package main

import (
	"net/http"

	"github.com/fpcoutinho/loja/routes"
)

func main() {
	routes.Init()
	http.ListenAndServe("localhost:8000", nil)
}
