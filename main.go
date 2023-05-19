package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectDB()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8000", nil)
}

// Produto é uma struct
type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// Index é a função que renderiza a página inicial
func index(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	selectProdutos, err := db.Query("select * from public.produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var nome, descricao string
		var preco float64
		var id, quantidade int
		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "index.html", produtos)
	defer db.Close()
}

// Função que conecta ao banco de dados
func connectDB() *sql.DB {
	conexao := "user=postgres dbname=alura-loja password=fpc050696 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
