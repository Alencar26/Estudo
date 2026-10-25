package database

import (
	"database/sql"

	"github.com/Alencar26/Estudo/GO/Clean-Architecture/internal/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?,?,?,?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	return err
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.db.QueryRow("SELECT count(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
