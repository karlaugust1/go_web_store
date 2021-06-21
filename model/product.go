package model

import (
	"go_web_store/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAll() []Product {
	db := db.ConectDatabase()

	selectAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConectDatabase()

	insertData, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConectDatabase()

	deleteData, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteData.Exec(id)

	defer db.Close()
}
