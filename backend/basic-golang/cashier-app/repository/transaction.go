package repository

import "strconv"

type TransactionRepository struct {
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository}
}

func (u *TransactionRepository) Pay(amount int) (int, error) {

	dataCSV, _ := u.cartItemRepository.db.Load("cart_items")
	total := 0
	for index, val := range dataCSV {
		if index != 0 {
			covertIntPrice, _ := strconv.Atoi(val[2])
			covertIntQuantity, _ := strconv.Atoi(val[3])
			total += covertIntPrice * covertIntQuantity
		}
	}

	return amount - total, nil // TODO: replace this
}
