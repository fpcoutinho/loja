package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/fpcoutinho/loja/models"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

// Index é a função que renderiza a página inicial
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.RetornaProdutos()
	temp.ExecuteTemplate(w, "index.html", produtos)
}

// New é a função que renderiza a página de cadastro de produtos
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new.html", nil)
}

// Insert é a função que insere um novo produto no banco de dados
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		p := r.FormValue("preco")
		q := r.FormValue("quantidade")

		preco, err := strconv.ParseFloat(p, 64)
		if err != nil {
			panic(err.Error())
		}
		quantidade, err := strconv.Atoi(q)
		if err != nil {
			panic(err.Error())
		}

		models.CriaProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
