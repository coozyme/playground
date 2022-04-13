package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)

	response := ProductListSuccessResponse{}
	response.Products = make([]Product, 0)

	products, err := api.productsRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

<<<<<<< HEAD
	for _, val := range products {
		response.Products = append(response.Products, Product{
			Name:     val.ProductName,
			Price:    val.Price,
			Category: val.Category,
		})
	}
	fmt.Println(products)

	encoder.Encode(ProductListSuccessResponse{Products: response.Products}) // TODO: replace this
=======
	fmt.Println(products)

	encoder.Encode(ProductListSuccessResponse{Products: []Product{}}) // TODO: replace this
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
}
