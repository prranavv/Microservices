package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "vk",
		Price: 2,
		SKU:   "ans-afd-ad",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
