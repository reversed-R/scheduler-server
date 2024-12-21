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
type Person struct {
	gorm.Model
	// Id      uint   `gorm:"primaryKey"`
	Name    string `gorm:"name"`
	Comment string `gorm:"comment"`
}

type Plan struct {
	gorm.Model
	// Id             uint `gorm:"primaryKey"`
	PersonId       uint `gorm:"person_id"`
	AvailabilityId uint `gorm:"availability_id"`
	TimeId         uint `gorm:"time_id"`
}

type Time struct {
	gorm.Model
	// Id   uint      `gorm:"primaryKey"`
	Time time.Time `gorm:"time"`
}

type Availability struct {
	gorm.Model
	// Id           uint   `gorm:"primaryKey"`
	Availability string `gorm:"availability"`
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func InitDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&Person{}, &Plan{}, &Time{}, &Availability{})

	db.Create(&Availability{Availability: "OK"})
	db.Create(&Availability{Availability: "NO"})
	db.Create(&Time{Time: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local)})
	db.Create(&Time{Time: time.Date(2024, 12, 25, 1, 30, 0, 0, time.Local)})
	CreatePerson(db, Person{Name: "John Smith", Comment: "Hello!"})
	CreatePerson(db, Person{Name: "Mary Smith", Comment: "Good Bye!"})
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
	var times []Time
	var availabilities []Availability
	var personalPlans []PersonalPlan
	availabilitiesMap := make(map[uint]string)
	timesMap := make(map[uint]time.Time)
	// var plans []Plan
	// personsResult := db.Find(&persons)
	db.Find(&persons)
	db.Find(&times)
	db.Find(&availabilities)

	fmt.Println("persons", persons)
	fmt.Println("times", times)
	fmt.Println("availabilities", availabilities)

	for _, a := range availabilities {
		availabilitiesMap[a.ID] = a.Availability
	}

	for _, t := range times {
		timesMap[t.ID] = t.Time
	}

	for _, person := range persons {
		plans := GetPlansByPersonId(db, person.ID)
		fmt.Println("plans", plans)
		plansMap := make(map[uint]uint) // plan // timeId -> AvailabilityId
		fmt.Println("plansMap", plansMap)
		for _, p := range plans {
			plansMap[p.TimeId] = p.AvailabilityId
			fmt.Println("plansMap", plansMap)
		}
		// var plans []Plan
		// db.Where("id = ?", person.Id).Find(&plans)
		// availabilityStrs := SliceMap(availabilities,
		// 	func(a Availability) string { return availabilitiesMap[a.Id] })

		availabilityStrs := SliceMap(times,
			func(t Time) string { return availabilitiesMap[plansMap[t.ID]] })
		// for _, a := range availabilities {
		// 	availabilityStrs = append(availabilityStrs, getAvailabilityById(availabilities, a.Id))
		// }

		personalPlans = append(personalPlans,
			PersonalPlan{
				Name:           person.Name,
				Comment:        person.Comment,
				Availabilities: availabilityStrs,
			})

		fmt.Println(availabilityStrs)
	}

	// personalPlans = SliceMap(persons,
	// 	func(person Person) PersonalPlan {
	// 		return SliceMap(GetPlansByPersonId(db, person.Id), func(plan Plan) PersonalPlan {
	// 			return PersonalPlan{
	// 				Name:    person.Name,
	// 				Comment: person.Comment,
	// 				Availabilities: SliceMap(availabilities,
	// 					func(a Availability) string { return availabilitiesMap[a.Id] }),
	// 			}
	// 		})
	// 	})

	// fmt.Println(Table{Times: times, PersonalPlans: personalPlans})

	return Table{Times: SliceMap(times, func(t Time) time.Time { return t.Time }), PersonalPlans: personalPlans}
}

func GetPlansByPersonId(db *gorm.DB, personId uint) []Plan {
	var plans []Plan
	db.Where("person_id = ?", personId).Find(&plans)

	return plans
}

// func getAvailabilityById(availabilities []Availability, id uint) string {
// 	for _, a := range availabilities {
// 		if a.Id == id {
// 			return a.Availability
// 		}
// 	}
// 	return ""
// }

// Update
func UpdatePersonName(db *gorm.DB, id uint, name string) {
	var oldPerson Person
	oldPerson.ID = id

	db.Model(&oldPerson).Update("Name", name)
}

func UpdatePersonComment(db *gorm.DB, id uint, comment string) {
	var oldPerson Person
	oldPerson.ID = id

	db.Model(&oldPerson).Update("Comment", comment)
}

// func UpdateByCode(db *gorm.DB, product *Product, code string) {
// 	// var oldAlbum Album
// 	// oldAlbum.ID = id
// 	// db.Model(oldAlbum).Updates(album)
// 	//
// 	// // Update - update product's price to 200
// 	// db.Model(&product).Update("Price", 200)
// 	// // Update - update multiple fields
// 	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
// 	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
// 	// fmt.Printf("product.Code=%s\n", product.Code)
// }
