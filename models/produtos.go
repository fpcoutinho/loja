package models

import (
	"github.com/fpcoutinho/loja/db"
	_ "github.com/lib/pq"
)

// Produto é uma struct
type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// CriaProduto é uma função que insere um novo produto no banco de dados
func CriaProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()
	insertProduto, err := db.Prepare("insert into public.produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertProduto.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func RetornaProdutos() []Produto {
	db := db.ConnectDB()
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
	defer db.Close()
	return produtos
}
