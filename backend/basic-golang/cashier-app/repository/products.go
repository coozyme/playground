package repository

import (
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
	record, err := u.db.Load("products")
	if err != nil {
		return nil, err
	}
	result := make([]Product, 0)
	for i := 1; i < len(record); i++ {
		Price, err := strconv.Atoi(record[i][2])
		if err != nil {
			return nil, err
		}

		product := Product{
			Category:    record[i][0],
			ProductName: record[i][1],
			Price:       Price,
		}
		result = append(result, product)
	}
	return result, err
	//return []Product{}, nil // TODO: replace this
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	product, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	return product, nil
	//	return []Product{}, nil // TODO: replace this
}
