// package classification of Product API
//
// # Documentation of Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// Consumes:
// - application/json
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AchrafMoreau/gRPC-in-Go/server/data"
	"github.com/gorilla/mux"
)

// Product represents a product handler
type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

// swagger:route GET /products products getAllProducts
// Returns a list of all products
// responses:
//   200: productsResponse
//   500: errorResponse
func (p *Product) GetAllProducts(resw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Get Request --")
	lp := data.GetProducts()
	err := lp.ToJson(resw)
	if err != nil {
		http.Error(resw, "Unabled to marshal Json", http.StatusInternalServerError)
	}
}

// swagger:route POST /products products addProduct
// Add a new product
// responses:
//   200: productResponse
//   400: errorResponse
//   500: errorResponse
func (p *Product) AddProducts(resw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Post Request --")

	prod := req.Context().Value(keyProduct{}).(data.Product)

	data.AddProduct(&prod)
	p.l.Printf("Add Product Successfully : %#v", prod)

}

// swagger:route PUT /products/{id} products updateProduct
// Update an existing product
// responses:
//   200: productResponse
//   400: errorResponse
//   404: errorResponse
func (p *Product) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStrig := vars["id"]
	id, iderr := strconv.Atoi(idStrig)
	if iderr != nil {
		http.Error(res, "Error Handling the Request", http.StatusBadRequest)
		return
	}

	p.l.Println("Hnadle PUT Request --", id)

	prod := req.Context().Value(keyProduct{}).(data.Product)

	er := data.UpdateProduct(id, &prod)
	if er == data.PRODUCT_NOT_FOUND_ERROR {
		http.Error(res, "Product Not Found", http.StatusNotFound)
		return
	}

	if er != nil {
		http.Error(res, "Error", http.StatusNotFound)
		return
	}
}

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product
// responses:
//   200: successResponse
//   400: errorResponse
//   404: errorResponse
func (p *Product) DeleteProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}

	p.l.Println("Hnadle Delete Request --", id)

	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}

	res.Write([]byte("Delete Succfully \n"))
}

type keyProduct struct{}

func (p Product) MiddlewareProduct(next http.Handler) http.Handler {
	p.l.Println("Middleware Called --")
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		prod := data.Product{}
		err := prod.FormatJson(req.Body)
		if err != nil {
			p.l.Println("[ERROR] Parrsing Error --", err)
			http.Error(res, "Error Parsing the Request", http.StatusBadRequest)
			return
		}

		err = prod.Validation()
		if err != nil{
			p.l.Println("[ERROR] Validation Error --", err)
			http.Error(res, fmt.Sprintf("Error Validate the Request Date", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(req.Context(), keyProduct{}, prod)
		request := req.WithContext(ctx)

		next.ServeHTTP(res, request)
	})
}

// You also need to define response models in your data package or here:

// swagger:response productsResponse
type productsResponseWrapper struct {
	// in:body
	Body []data.Product
}

// swagger:response productResponse
type productResponseWrapper struct {
	// in:body
	Body data.Product
}

// swagger:response errorResponse
type errorResponseWrapper struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response successResponse
type successResponseWrapper struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}
