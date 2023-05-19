package routes

import (
	"net/http"

	"github.com/fpcoutinho/loja/controllers"
)

func Init() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
