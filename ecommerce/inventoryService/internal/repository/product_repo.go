package repository

import (
	"github.com/jmoiron/sqlx"
	"inventoryService/internal/domain"
)

// 9. Тут уже идет выполнение SQL-запроса в Postgres
type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) domain.ProductRepository {
	return &productRepo{db}
}

// 9.1 После чего через Handler возвращается наш результат(p) в JSON-формате обратно клиенту(handler.go)
func (r *productRepo) Create(p *domain.Product) error {
	query := `INSERT INTO products (name, description, price, stock)
			  VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, p.Name, p.Description, p.Price, p.Stock).Scan(&p.ID)
}

func (r *productRepo) GetByID(id int) (*domain.Product, error) {
	var p domain.Product
	err := r.db.Get(&p, "SELECT * FROM products WHERE id=$1", id)
	return &p, err
}

func (r *productRepo) Update(p *domain.Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, stock=$4 WHERE id=$6`
	_, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.ID)
	return err
}

func (r *productRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}

func (r *productRepo) List() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Select(&products, "SELECT * FROM products")
	return products, err
}
