package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectDatabase() *sql.DB {
	conection := "user=postgres dbname=go_web_store password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectDatabase()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectDatabase()

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

	temp.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}
