package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/mrpsousa/go-hexagonal/adapters/db"
	"github.com/mrpsousa/go-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}
