package handlers

import (
	"net/http"
	"project/data"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handle Delete Product", id)
	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Produc tnot found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}
