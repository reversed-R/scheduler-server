package internal

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// for front
type Table struct {
	// Title         string         `json:"title"`
	// Description   string         `json:"description"`
	Times         []time.Time    `json:"times"`
	PersonalPlans []PersonalPlan `json:"personalPlans"`
}

type PersonalPlan struct {
	Name           string   `json:"name"`
	Comment        string   `json:"comment"`
	Availabilities []string `json:"availabilities"`
}

// for DB tables
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Person struct {
	gorm.Model
	Id      uint `gorm:"primaryKey"`
	Name    string
	Comment string
}

type Plan struct {
	gorm.Model
	PersonId       uint
	AvailabilityId uint
	TimeId         uint
}

type Time struct {
	gorm.Model
	Id   uint `gorm:"primaryKey"`
	Time time.Time
}

type Availability struct {
	gorm.Model
	Id           uint `gorm:"primaryKey"`
	Availability string
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func InitDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&Person{}, &Plan{}, &Time{}, &Availability{})

	db.Create(&Availability{Id: 1, Availability: "OK"})
	db.Create(&Availability{Id: 2, Availability: "NO"})
	db.Create(&Time{Id: 1, Time: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local)})
	db.Create(&Time{Id: 2, Time: time.Date(2024, 12, 25, 1, 0, 0, 0, time.Local)})
	CreatePerson(db, Person{Id: 1, Name: "John Smith", Comment: "Hello!"})
	CreatePerson(db, Person{Id: 2, Name: "Mary Smith", Comment: "Good Bye!"})
	CreatePlan(db, Plan{PersonId: 1, TimeId: 1, AvailabilityId: 2})
	CreatePlan(db, Plan{PersonId: 1, TimeId: 2, AvailabilityId: 1})
	CreatePlan(db, Plan{PersonId: 2, TimeId: 1, AvailabilityId: 2})
	CreatePlan(db, Plan{PersonId: 2, TimeId: 2, AvailabilityId: 2})
}

// Create
func CreatePerson(db *gorm.DB, person Person) {
	db.Create(&person)
}

func CreatePlan(db *gorm.DB, plan Plan) {
	db.Create(&plan)
}

// Read
// func ReadProductFirstByCode(db gorm.DB, product *Product, code string) {
// 	db.First(&product, code)
// }

func GetTable(db *gorm.DB) Table {
	var persons []Person
	var times []time.Time
	var availabilities []Availability
	var personalPlans []PersonalPlan
	// var plans []Plan
	// personsResult := db.Find(&persons)
	db.Find(&persons)
	db.Find(&times)
	db.Find(&availabilities)

	for _, person := range persons {
		var plans []Plan
		db.Where("id = ?", person.Id).Find(&plans)
		var availabilityStrs []string

		for _, a := range availabilities {
			availabilityStrs = append(availabilityStrs, getAvailabilityById(availabilities, a.Id))
		}

		personalPlans = append(personalPlans,
			PersonalPlan{
				Name:           person.Name,
				Comment:        person.Comment,
				Availabilities: availabilityStrs,
			})

		fmt.Println(availabilityStrs)
	}

	fmt.Println(Table{Times: times, PersonalPlans: personalPlans})

	return Table{Times: times, PersonalPlans: personalPlans}
}

func getAvailabilityById(availabilities []Availability, id uint) string {
	for _, a := range availabilities {
		if a.Id == id {
			return a.Availability
		}
	}
	return ""
}

// Update
func UpdatePersonName(db *gorm.DB, id uint, name string) {
	var oldPerson Person
	oldPerson.Id = id

	db.Model(&oldPerson).Update("Name", name)
}

func UpdatePersonComment(db *gorm.DB, id uint, comment string) {
	var oldPerson Person
	oldPerson.Id = id

	db.Model(&oldPerson).Update("Comment", comment)
}

func UpdateByCode(db *gorm.DB, product *Product, code string) {
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
