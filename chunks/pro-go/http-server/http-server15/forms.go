package main

import (
	"net/http"
	"strconv"
)

func ProcessFormData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		index, _ := strconv.Atoi(r.PostFormValue("index"))
		p := Product{}
		p.Name = r.PostFormValue("name")
		p.Category = r.PostFormValue("category")
		p.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 64)
		Printfln("From ProcessFormData(): index - %v", index)
		Printfln("From ProcessFormData(): p.Name - %s", p.Name)
		Printfln("From ProcessFormData(): p.Category - %s", p.Category)
		Printfln("From ProcessFormData(): p.Price - %.2f", p.Price)
		Products[index] = p
	}
	http.Redirect(w, r, "/templates", http.StatusTemporaryRedirect)
}

func init() {
	http.HandleFunc("/forms/edit", ProcessFormData)
}
