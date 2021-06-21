package controllers

import (
	"go_web_store/model"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.FindAll()
	temp.ExecuteTemplate(w, "Index", products)
}
