package internal

import (
	// "fmt"
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
	Name    string `gorm:"name"`
	Comment string `gorm:"comment"`
}

type Plan struct {
	gorm.Model
	PersonId       uint `gorm:"person_id"`
	AvailabilityId uint `gorm:"availability_id"`
	TimeId         uint `gorm:"time_id"`
}

type Time struct {
	gorm.Model
	Time time.Time `gorm:"time"`
}

type Availability struct {
	gorm.Model
	Availability string `gorm:"availability"`
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}

func InitDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&Person{}, &Plan{}, &Time{}, &Availability{})

	// db.Create(&Availability{Availability: "OK"})
	// db.Create(&Availability{Availability: "NO"})
	// db.Create(&Time{Time: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local)})
	// db.Create(&Time{Time: time.Date(2024, 12, 25, 1, 30, 0, 0, time.Local)})
	// CreatePerson(db, Person{Name: "John Smith", Comment: "Hello!"})
	// CreatePerson(db, Person{Name: "Mary Smith", Comment: "Good Bye!"})
	// CreatePlan(db, Plan{PersonId: 1, TimeId: 1, AvailabilityId: 2})
	// CreatePlan(db, Plan{PersonId: 1, TimeId: 2, AvailabilityId: 1})
	// CreatePlan(db, Plan{PersonId: 2, TimeId: 1, AvailabilityId: 2})
	// CreatePlan(db, Plan{PersonId: 2, TimeId: 2, AvailabilityId: 2})
}

// Create
func CreatePerson(db *gorm.DB, person Person) (*gorm.DB, Person) {
	result := db.Create(&person)
	return result, person
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

	db.Find(&persons)
	db.Find(&times)
	db.Find(&availabilities)

	// fmt.Println("persons", persons)
	// fmt.Println("times", times)
	// fmt.Println("availabilities", availabilities)

	for _, a := range availabilities {
		availabilitiesMap[a.ID] = a.Availability
	}

	for _, t := range times {
		timesMap[t.ID] = t.Time
	}

	for _, person := range persons {
		plans := GetPlansByPersonId(db, person.ID)
		// fmt.Println("plans", plans)
		plansMap := make(map[uint]uint) // plan // timeId -> AvailabilityId
		// fmt.Println("plansMap", plansMap)
		for _, p := range plans {
			plansMap[p.TimeId] = p.AvailabilityId
			// fmt.Println("plansMap", plansMap)
		}

		availabilityStrs := SliceMap(times,
			func(t Time) string { return availabilitiesMap[plansMap[t.ID]] })

		personalPlans = append(personalPlans,
			PersonalPlan{
				Name:           person.Name,
				Comment:        person.Comment,
				Availabilities: availabilityStrs,
			})
	}

	return Table{Times: SliceMap(times, func(t Time) time.Time { return t.Time }), PersonalPlans: personalPlans}
}

func GetPlansByPersonId(db *gorm.DB, personId uint) []Plan {
	var plans []Plan
	db.Where("person_id = ?", personId).Find(&plans)

	return plans
}

// Update
func UpdatePerson(db *gorm.DB, id uint, person Person) (*gorm.DB, Person) {
	var oldPerson Person
	oldPerson.ID = id

	result := db.Model(&oldPerson).Updates(person)

	return result, oldPerson
}

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
