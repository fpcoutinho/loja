package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	conexao := fmt.Sprintf(`user=%s dbname=%s password=%s host=localhost sslmode=disable`, os.Getenv("DB_USR"), os.Getenv("DB_NAME"), os.Getenv("DB_PWD"))

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
