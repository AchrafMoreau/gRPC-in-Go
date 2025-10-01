package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"slices"
	"time"
	"github.com/go-playground/validator/v10"
)
// Product represents a product in the system
//
// swagger:model
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required,desc"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var PRODUCT_NOT_FOUND_ERROR = fmt.Errorf("Product Not Found !")

// Products represents a list of products
//
// swagger:model
type Products []*Product

// ===========================
// VALIDATION FOR CRUD FUNCTIONS
// ===========================
func (p *Product) Validation() error{
	validator := validator.New()
	validator.RegisterValidation("sku", validationSKU)
	validator.RegisterValidation("desc", validateDescription)
	return validator.Struct(p)
}

func validateDescription(fl validator.FieldLevel) bool{
	if len(fl.Field().String()) < 10 {
		return false
	}
	return true
}

func validationSKU(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	matches := regex.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1{
		return false
	}

	return true
}
// ===========================
// UTILITY FOR CRUD FUNCTIONS
// ===========================

func (p *Product)FormatJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func findProduct(id int) (*Product, int, error){
	for pos, prod := range ProductList{
		if prod.ID == id{
			return prod, pos, nil
		}
	}
	return nil, 0, PRODUCT_NOT_FOUND_ERROR
}

func getNextProductId() int{
	ls := ProductList[len(ProductList) - 1];
	return ls.ID + 1
}

// ===========================
// CRUD FUNCTIONS ----
// ===========================

func GetProducts() Products {
	return ProductList
}

func AddProduct(prod *Product){
	prod.ID = getNextProductId();
	ProductList = append(ProductList, prod)
}

func UpdateProduct(id int, prod *Product) error {
	_, pos, err := findProduct(id)
	if err != nil{
		return err
	}

	prod.ID = id
	ProductList[pos] = prod
	return nil
}

func DeleteProduct(id int) error{
	_, pos, err := findProduct(id)
	if err != nil{
		return err
	}

	ProductList = slices.Delete(ProductList, pos, pos + 1)
	return nil
}


// ===========================
// EXAMPLE OF COLLECTION
// ===========================

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy Milk Coffee",
		Price:       2.25,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espersso",
		Description: "Short and Strong Caffee whitout Milk",
		Price:       1.99,
		SKU:         "jhk123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
