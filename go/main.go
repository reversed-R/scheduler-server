package main

import (
	// "database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Album struct {
// 	gorm.Model
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Printf("main---->\n")

	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	fmt.Printf("product.Code=%s\n", product.Code)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Printf("product.Code=%s\n", product.Code)

	// Delete - delete product
	db.Delete(&product, 1)
	fmt.Printf("product.Code=%s\n", product.Code)

	fmt.Printf("<----main\n")
}

// Create
func CreateProduct(db gorm.DB, product Product) {
	db.Create(&product)
}

// Read
func ReadProductFirstByCode(db gorm.DB, product *Product, code string) {
	db.First(&product, code)
}

// Update
func UpdateByCode(db gorm.DB, product *Product, code string) {
	// var oldAlbum Album
	// oldAlbum.ID = id
	// db.Model(oldAlbum).Updates(album)
	//
	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// fmt.Printf("product.Code=%s\n", product.Code)

}

// // Delete
// func Delete(album Album) {
//         db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//         if err != nil {
//                 panic("failed to connect database")
//         }
//
//         db.Delete(&album, 1)
// }

// type Product struct {
//      gorm.Model
//      Code  string
//      Price uint
// }
//
// func test() {
//      db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//      if err != nil {
//              panic("failed to connect database")
//      }
//
//      // Migrate the schema
//      db.AutoMigrate(&Product{})
//
//      // Create
//      db.Create(&Product{Code: "D42", Price: 100})
//
//      // Read
//      var product Product
//      db.First(&product, 1)
//      db.First(&product, "code = ?", "D42")
//
//      // Update
//      db.Model(&product).Update("Price", 200)
//      // Update
//      db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
//      db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
//
//      // Delete
//      db.Delete(&product, 1)
// }
