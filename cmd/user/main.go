package main

import (
	"fmt"

	dbp "mirror/pkg/dbplugin"
	"mirror/pkg/gomysql"
)

type Product struct {
	dbp.M
	Code  string
	Price uint
}

func main() {
	err := gomysql.Init(&gomysql.Config{
		Addr: "127.0.0.1",
		User: "root",
		Pass: "123456",
		DB:   "lucky",
	})

	if err != nil {
		fmt.Println(err)
	}

	db := gomysql.Use("lucky")

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
