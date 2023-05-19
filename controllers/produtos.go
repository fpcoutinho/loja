package controllers

import (
	"html/template"
	"net/http"

	"github.com/fpcoutinho/loja/models"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

// Index é a função que renderiza a página inicial
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.RetornaProdutos()
	temp.ExecuteTemplate(w, "index.html", produtos)
}
