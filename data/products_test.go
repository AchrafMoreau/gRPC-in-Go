package data

import (
	"testing"
)

func TestValidaiton(t *testing.T) {
	prod := &Product{
		Name: "Achraf",
		Price: 20.00,
		SKU: "ghi-smaial-yassin",
	}

	err := prod.validation()
	if err != nil{
		t.Fatal(err)
	}

}
