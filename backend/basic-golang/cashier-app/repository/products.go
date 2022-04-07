package repository

import (
	// "strconv"

	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type ProductRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) LoadOrCreate() ([]Product, error) {
	return []Product{}, nil // TODO: replace this
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	data, _ := u.db.Load("products")
	var products []Product

	for index, val := range data {
		if index != 0 {
			convertToInt, _ := strconv.Atoi(val[2])
			products = append(products, Product{
				Category:    val[0],
				ProductName: val[1],
				Price:       convertToInt,
			})
		}
	}
	return products, nil // TODO: replace this
}
