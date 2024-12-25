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
	Times     []time.Time `json:"times"`
	UserPlans []UserPlan  `json:"userPlans"`
}

type UserPlan struct {
	Name           string   `json:"name"`
	Comment        string   `json:"comment"`
	Availabilities []string `json:"availabilities"`
}

// for DB tables
type User struct {
	gorm.Model
	Name    string `gorm:"name"`
	Comment string `gorm:"comment"`
}

type Plan struct {
	gorm.Model
	UserId         uint `gorm:"user_id"`
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
	db.AutoMigrate(&User{}, &Plan{}, &Time{}, &Availability{})

	// db.Create(&Availability{Availability: "OK"})
	// db.Create(&Availability{Availability: "NO"})
	// db.Create(&Time{Time: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local)})
	// db.Create(&Time{Time: time.Date(2024, 12, 25, 1, 30, 0, 0, time.Local)})
	// CreateUser(db, User{Name: "John Smith", Comment: "Hello!"})
	// CreateUser(db, User{Name: "Mary Smith", Comment: "Good Bye!"})
	// CreatePlan(db, Plan{UserId: 1, TimeId: 1, AvailabilityId: 2})
	// CreatePlan(db, Plan{UserId: 1, TimeId: 2, AvailabilityId: 1})
	// CreatePlan(db, Plan{UserId: 2, TimeId: 1, AvailabilityId: 2})
	// CreatePlan(db, Plan{UserId: 2, TimeId: 2, AvailabilityId: 2})
}

// Create
func CreateUser(db *gorm.DB, user User) (*gorm.DB, User) {
	result := db.Create(&user)
	return result, user
}

func CreatePlan(db *gorm.DB, plan Plan) {
	db.Create(&plan)
}

// Read
// func ReadProductFirstByCode(db gorm.DB, product *Product, code string) {
// 	db.First(&product, code)
// }

func GetTable(db *gorm.DB) Table {
	var users []User
	var times []Time
	var availabilities []Availability
	var userPlans []UserPlan
	availabilitiesMap := make(map[uint]string)
	timesMap := make(map[uint]time.Time)

	db.Find(&users)
	db.Find(&times)
	db.Find(&availabilities)

	fmt.Println("users", users)
	fmt.Println("times", times)
	fmt.Println("availabilities", availabilities)

	for _, a := range availabilities {
		availabilitiesMap[a.ID] = a.Availability
	}

	for _, t := range times {
		timesMap[t.ID] = t.Time
	}

	for _, user := range users {
		plans := GetPlansByUserId(db, user.ID)
		// fmt.Println("plans", plans)
		plansMap := make(map[uint]uint) // plan // timeId -> AvailabilityId
		// fmt.Println("plansMap", plansMap)
		for _, p := range plans {
			plansMap[p.TimeId] = p.AvailabilityId
			fmt.Println("plansMap", plansMap)
		}

		availabilityStrs := SliceMap(times,
			func(t Time) string { return availabilitiesMap[plansMap[t.ID]] })

		userPlans = append(userPlans,
			UserPlan{
				Name:           user.Name,
				Comment:        user.Comment,
				Availabilities: availabilityStrs,
			})
	}

	return Table{Times: SliceMap(times, func(t Time) time.Time { return t.Time }), UserPlans: userPlans}
}

func GetPlansByUserId(db *gorm.DB, userId uint) []Plan {
	var plans []Plan
	db.Where("user_id = ?", userId).Find(&plans)

	return plans
}

// Update
func UpdateUser(db *gorm.DB, id uint, user User) (*gorm.DB, User) {
	var oldUser User
	oldUser.ID = id

	result := db.Model(&oldUser).Updates(user)

	return result, oldUser
}

func UpdateUserName(db *gorm.DB, id uint, name string) {
	var oldUser User
	oldUser.ID = id

	db.Model(&oldUser).Update("Name", name)
}

func UpdateUserComment(db *gorm.DB, id uint, comment string) {
	var oldUser User
	oldUser.ID = id

	db.Model(&oldUser).Update("Comment", comment)
}
