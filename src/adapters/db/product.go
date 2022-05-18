package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sergio/go-hexagonal/application/product"
	appProduct "github.com/sergio/go-hexagonal/application/product"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}

func (p *ProductDB) List() ([]appProduct.IProduct, error) {
	var products []appProduct.IProduct

	stmt, err := p.db.Prepare("SELECT id, name, status, price FROM products;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product product.Product
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Status,
			&product.Price,
		)
		products = append(products, &product)
	}

	return products, nil
}

func (p *ProductDB) Get(id string) (appProduct.IProduct, error) {
	var product product.Product

	stmt, err := p.db.Prepare("SELECT id, name, status, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Status,
		&product.Price,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) Save(product appProduct.IProduct) error {
	var rows int
	p.db.QueryRow("SELECT COUNT(id) FROM products WHERE id = ?;", product.GetID()).Scan(&rows)
	if rows != 0 {
		return p.update(product)
	}

	return p.create(product)
}

func (p *ProductDB) update(product appProduct.IProduct) error {
	stmt, err := p.db.Prepare(`
		UPDATE products
		SET 
			name = ?
			, status = ?
			, price = ?
		WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		product.GetName(),
		product.GetStatus(),
		product.GetPrice(),
		product.GetID(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDB) create(product appProduct.IProduct) error {
	stmt, err := p.db.Prepare("INSERT INTO products (id, name, status, price) VALUES (?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetStatus(), product.GetPrice())
	if err != nil {
		return err
	}
	return nil
}
