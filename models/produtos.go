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
