package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	//slu is of format avc=ava-sadfs
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

func AddProd(p *Product) {
	p.ID = getnextID()
	productlist = append(productlist, p)
}

func getnextID() int {
	lp := productlist[len(productlist)-1]
	return lp.ID + 1
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func findIndexByProductID(id int) int {
	for i, p := range productlist {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var ErrProductNotFound = fmt.Errorf("Prodtct not found")

func findproduct(id int) (*Product, int, error) {
	for i, p := range productlist {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findproduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productlist[pos] = p
	return nil
}

func GetProducts() Products {
	return productlist
}

var productlist = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "asdlk",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "djsk",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productlist = append(productlist[:i], productlist[i+1])

	return nil
}
