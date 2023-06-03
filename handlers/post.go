package handlers

import (
	"net/http"
	"project/data"
)

func (p *Products) Addproduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")
	prod := r.Context().Value(Keyproduct{}).(data.Product)

	data.AddProd(&prod)
}
