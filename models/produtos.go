package models

import (
	"github.com/fpcoutinho/loja/db"
	_ "github.com/lib/pq"
)

// Produto é uma struct
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// GetProduto é uma função que retorna um produto do banco de dados
func GetProduto(id string) Produto {
	db := db.ConnectDB()
	selectProduto, err := db.Query("select * from public.produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.ID = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}
	defer db.Close()
	return produto
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

// DeletaProduto é uma função que deleta um produto do banco de dados
func DeletaProduto(id string) {
	db := db.ConnectDB()
	deleteProduto, err := db.Prepare("delete from public.produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteProduto.Exec(id)
	defer db.Close()
}

// UpdateProduto é uma função que atualiza um produto no banco de dados
func UpdateProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()
	updateProduto, err := db.Prepare("update public.produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduto.Exec(nome, descricao, preco, quantidade, id)
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
		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
