package model

import (
	"go_web_store/db"
)

type Product struct {
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
			p.Name = name
			p.Description = description
			p.Price = price
			p.Quantity = quantity

			products = append(products, p)
		}
	}

	defer db.Close()

	return products
}
