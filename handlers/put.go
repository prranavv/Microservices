package handlers

import (
	"net/http"
	"project/data"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) Updateproducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert ID", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT product")
	prod := r.Context().Value(Keyproduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
