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

// Delete é a função que deleta um produto do banco de dados
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Edit é a função que renderiza a página de edição de produtos
func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.GetProduto(idProduto)
	temp.ExecuteTemplate(w, "edit.html", produto)
}

// Update é a função que atualiza um produto no banco de dados
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		p := r.FormValue("preco")
		q := r.FormValue("quantidade")

		idProduto, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}
		preco, err := strconv.ParseFloat(p, 64)
		if err != nil {
			panic(err.Error())
		}
		quantidade, err := strconv.Atoi(q)
		if err != nil {
			panic(err.Error())
		}

		models.UpdateProduto(idProduto, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
