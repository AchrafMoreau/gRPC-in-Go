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

	err := prod.Validation()
	if err != nil{
		t.Fatal(err)
	}

}
