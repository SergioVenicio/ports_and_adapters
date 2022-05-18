package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sergio/go-hexagonal/adapters/db"
	"github.com/sergio/go-hexagonal/application/product"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")

	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
		"id" VARCHAR(255),
		"name" VARCHAR(255),
		"status" VARCHAR(255),
		"price" FLOAT
 	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	id := "123"
	name := "Test"
	status := "enabled"
	price := 1.99
	insert := `INSERT INTO products (id, name, status, price) VALUES (?, ?, ?, ?);`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec(id, name, status, price)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product, err := productDB.Get("123")

	require.Nil(t, err)
	require.Equal(t, "Test", product.GetName())
	require.Equal(t, "enabled", product.GetStatus())
	require.Equal(t, 1.99, product.GetPrice())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDB(DB)

	newProduct := product.NewProduct("Test", 1.99, product.ENABLED)
	err := productDB.Save(newProduct)

	require.Nil(t, err)
	require.Equal(t, newProduct.GetName(), "Test")
	require.Equal(t, newProduct.GetStatus(), "enabled")
	require.Equal(t, newProduct.GetPrice(), 1.99)
}

func TestProductDB_SaveExistentProduct(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDB(DB)

	newProduct := product.NewProduct("Test", 1.99, product.ENABLED)
	productDB.Save(newProduct)

	newProduct.Disable()
	err := productDB.Save(newProduct)

	require.Nil(t, err)
	require.Equal(t, newProduct.GetName(), "Test")
	require.Equal(t, newProduct.GetStatus(), "disabled")
	require.Equal(t, newProduct.GetPrice(), 1.99)
}
