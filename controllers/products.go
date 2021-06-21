package controllers

import (
	"go_web_store/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.FindAll()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço")
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantidade")
		}

		model.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", 301)
}
