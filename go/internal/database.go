package internal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Migrate the schema
func Migrate() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Album{})
}

// Create
func Create(album Album) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&album)
}

// Read
func Read() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var album Album
	db.First(&album, 1)
}

// Update
func UpdateById(id string, album Album) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var oldAlbum Album
	oldAlbum.ID = id
	db.Model(oldAlbum).Updates(album)
}

// Delete
func Delete(album Album) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Delete(&album, 1)
}

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }
//
// func test() {
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
//
// 	// Migrate the schema
// 	db.AutoMigrate(&Product{})
//
// 	// Create
// 	db.Create(&Product{Code: "D42", Price: 100})
//
// 	// Read
// 	var product Product
// 	db.First(&product, 1)
// 	db.First(&product, "code = ?", "D42")
//
// 	// Update
// 	db.Model(&product).Update("Price", 200)
// 	// Update
// 	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
// 	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
//
// 	// Delete
// 	db.Delete(&product, 1)
// }
