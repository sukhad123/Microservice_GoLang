package products

import (
	"fmt"
	"net/http"

	"github.com/learn_go/internal/json"
)

// this will go on server folder
//class
/**
class Sukhad{
name String
}*/
type handler struct {
	service Service
}

/**Sukhad() constructor*/
//Constructor function
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// function
// method inside constructor
func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	//come here when we make request
	fmt.Println("I am called for the api products request")
	//return something in json
	//products := []string{"hello", "whats up"}
	//Fetch products from db
	products, err := h.service.ListProducts(r.Context())

	//throw error from an api
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//set ok
	//json.Write(w, http.StatusOK, products)
	json.Write(w, http.StatusOK, products)

}

// method inside constructor
func (h *handler) addProducts(w http.ResponseWriter, r *http.Request) {
	//come here when we make request
	fmt.Println("I am called for the post api products request")
	//return something in json
	// Decode JSON request body
	// var req struct {
	// 	Name     string
	// 	Price    float64
	// 	Quantity int
	// }

	// params := products.CreateProductParams{
	// 	Name:     pgtype.Text{String: req.Name, Valid: true},
	// 	Price:    req.Price,
	// 	Quantity: req.Quantity,
	// }

	//set ok
	//json.Write(w, http.StatusOK, products)
}
