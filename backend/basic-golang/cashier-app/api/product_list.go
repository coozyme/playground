package api

import (
	"encoding/json"
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
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)
	// _, err := api.AuthMiddleWare(w, req)
	// if err != nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	encoder.Encode(ProductListErrorResponse{Error: err.Error()})
	// 	return
	// }

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
=======
	// fmt.Println(products)
>>>>>>> 264ddc6dc02457d5a01ce89929fbd0225aab642b

	encoder.Encode(ProductListSuccessResponse{Products: []Product{}}) // TODO: replace this
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
}
