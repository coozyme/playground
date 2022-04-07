package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type CartItemRepository struct {
	db db.DB
}

func NewCartItemRepository(db db.DB) CartItemRepository {
	return CartItemRepository{db}
}

func (u *CartItemRepository) LoadOrCreate() ([]CartItem, error) {
	records, err := u.db.Load("cart_items")
	if err != nil {
		records = [][]string{
			{"category", "product_name", "price", "quantity"},
		}
		if err := u.db.Save("cart_items", records); err != nil {
			return nil, err
		}
	}

	result := make([]CartItem, 0)
	for i := 1; i < len(records); i++ {
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		qty, err := strconv.Atoi(records[i][3])
		if err != nil {
			return nil, err
		}

		cartItem := CartItem{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
			Quantity:    qty,
		}
		result = append(result, cartItem)
	}

	return result, nil
}

func (u *CartItemRepository) Save(cartItems []CartItem) error {
	records := [][]string{
		{"category", "product_name", "price", "quantity"},
	}
	for i := 0; i < len(cartItems); i++ {
		records = append(records, []string{
			cartItems[i].Category,
			cartItems[i].ProductName,
			strconv.Itoa(cartItems[i].Price),
			strconv.Itoa(cartItems[i].Quantity),
		})
	}
	return u.db.Save("cart_items", records)
}

func (u *CartItemRepository) SelectAll() ([]CartItem, error) {
	dataCSV, _ := u.db.Load("cart_items")
	var carItems []CartItem
	for index, val := range dataCSV {
		if index != 0 {
			covertIntPrice, _ := strconv.Atoi(val[2])
			covertIntQuantity, _ := strconv.Atoi(val[3])
			carItems = append(carItems, CartItem{
				Category:    val[0],
				ProductName: val[1],
				Price:       covertIntPrice,
				Quantity:    covertIntQuantity,
			})
		}
	}

	return carItems, nil // TODO: replace this
}

func (u *CartItemRepository) Add(product Product) error {
	dataCSV, err := u.db.Load("cart_items")
	if err != nil {
		records := [][]string{
			{"category", "product_name", "price", "quantity"},
		}
		if err := u.db.Save("cart_items", records); err != nil {
			return err
		}
		return nil
	}

	isFound := false
	for index, val := range dataCSV {
		if index != 0 {

			if val[1] == product.ProductName {
				parseInt, _ := strconv.Atoi(val[3])
				dataCSV[index][3] = strconv.Itoa(parseInt + 1)
				isFound = true
				break
			}
		}
	}

	if isFound == false {
		parseString := strconv.Itoa(product.Price)
		dataCSV = append(dataCSV, []string{
			product.Category, product.ProductName, parseString, "1",
		})
	}

	u.db.Save("cart_items", dataCSV)

	return nil // TODO: replace this
}

func (u *CartItemRepository) ResetCartItems() error {
	u.db.Delete("cart_items")
	cartItem := [][]string{{"category", "product_name", "price", "quantity"}}

	u.db.Save("cart_items", cartItem)
	return nil // TODO: replace this
}

func (u *CartItemRepository) TotalPrice() (int, error) {
	dataCSV, _ := u.db.Load("cart_items")
	total := 0
	for index, val := range dataCSV {
		if index != 0 {
			covertIntPrice, _ := strconv.Atoi(val[2])
			covertIntQuantity, _ := strconv.Atoi(val[3])

			total += covertIntPrice * covertIntQuantity
		}
	}
	return total, nil // TODO: replace this
}
