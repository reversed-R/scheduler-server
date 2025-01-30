package internal

import (
	// "fmt"
	// "github.com/thoas/go-funk"
	// "golang.org/x/exp/slices"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "sort"
	// "time"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}

func InitDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&Room{}, &User{}, &Plan{})
	// db.AutoMigrate(&Room{}, &User{}, &Plan{}, &Availability{})
	// db.AutoMigrate(&Room{}, &User{}, &Plan{}, &Time{}, &Availability{})

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
func CreateRoom(db *gorm.DB, room Room) (*gorm.DB, Room) {
	result := db.Create(&room)
	return result, room
}

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

func GetRoomAllInfo(db *gorm.DB, roomId uint) (RoomAllInfoJSON, error) {
	// var users []User
	// var times []Time
	// var availabilities []Availability
	// var availabilities []string
	var userJSONs []UserJSON
	// availabilitiesMap := make(map[uint]string)
	// timesMap := make(map[uint]time.Time)

	room, _ := GetRoom(db, roomId)
	users, _ := GetUsersByRoomId(db, roomId)
	// db.Find(&users)
	// db.Find(&times)
	// db.Find(&availabilities)

	// fmt.Println("users", users)
	// fmt.Println("times", times)
	// fmt.Println("availabilities", availabilities)

	// for _, a := range availabilities {
	// 	availabilitiesMap[a.ID] = a.Availability
	// }

	// for _, t := range times {
	// 	timesMap[t.ID] = t.Time
	// }

	// for _, user := range users {
	// 	plans, _ := GetPlansByUserId(db, user.ID)
	// 	// fmt.Println("plans", plans)
	// 	plansMap := make(map[uint]uint) // plan // timeId -> AvailabilityId
	// 	// fmt.Println("plansMap", plansMap)
	// 	for _, p := range plans {
	// 		plansMap[p.TimeId] = p.AvailabilityId
	// 		fmt.Println("plansMap", plansMap)
	// 	}
	//
	// 	availabilityStrs := SliceMap(times,
	// 		func(t Time) string { return availabilitiesMap[plansMap[t.ID]] })
	//
	// 	userPlans = append(userPlans,
	// 		UserJSON{
	// 			Name:           user.Name,
	// 			Comment:        user.Comment,
	// 			Availabilities: availabilityStrs,
	// 		})
	// }

	for _, user := range users {
		plans, _ := GetPlansByUserId(db, user.ID)

		// sort.Slice(plans, func(i, j int) bool { return plans[i].TimeId < plans[j].TimeId })

		var availabilities []string
		// var availabilities [room.DayLength * room.DayPatternLength]string
		for i := 0; i < int(room.DayLength*room.DayPatternLength); i++ {
			// plan := funk.Find(plans, func(plan Plan) bool { return int(plan.TimeId) == i })
			availabilities = append(availabilities, "")

			for _, p := range plans {
				if int(p.TimeId) == i {
					availabilities[i] = p.Availability
					break
				}
			}

			// availabilities[i] = plan.Availability
		}

		userJSONs = append(userJSONs, UserJSON{
			Name:           user.Name,
			Comment:        user.Comment,
			Availabilities: availabilities,
		})
	}

	return RoomAllInfoJSON{
			RoomName:        room.Name,
			RoomDescription: room.Description,
			BeginTime: TimeJSON{
				Year:  room.BeginTime.Year(),
				Month: room.BeginTime.Month(),
				Day:   room.BeginTime.Day(),
				Hour:  room.BeginTime.Hour(),
				Min:   room.BeginTime.Minute(),
			},
			DayLength:        room.DayLength,
			DayPattern:       room.DayPattern,
			DayPatternLength: room.DayPatternLength,
			Users:            userJSONs,
		},
		nil

	// return RoomAllInfo{
	// 		RoomName:        room.Name,
	// 		RoomDescription: room.Description,
	// 		Times:           SliceMap(times, func(t Time) time.Time { return t.Time }),
	// 		UserPlans:       userPlans},
	// 	nil
}

func GetRoom(db *gorm.DB, roomId uint) (Room, error) {
	var room Room
	result := db.Where("id = ?", roomId).Find(&room)

	return room, result.Error
}

func GetUsersByRoomId(db *gorm.DB, roomId uint) ([]User, error) {
	var users []User
	result := db.Where("room_id = ?", roomId).Find(&users)

	return users, result.Error
}

func GetPlansByUserId(db *gorm.DB, userId uint) ([]Plan, error) {
	var plans []Plan
	result := db.Where("user_id = ?", userId).Find(&plans)

	return plans, result.Error
}

// Update
func UpdateUser(db *gorm.DB, id uint, user User) (*gorm.DB, User) {
	var oldUser User
	oldUser.ID = id

	result := db.Model(&oldUser).Updates(user)

	return result, oldUser
}

// func UpdateUserName(db *gorm.DB, id uint, name string) {
// 	var oldUser User
// 	oldUser.ID = id
//
// 	db.Model(&oldUser).Update("Name", name)
// }
//
// func UpdateUserComment(db *gorm.DB, id uint, comment string) {
// 	var oldUser User
// 	oldUser.ID = id
//
// 	db.Model(&oldUser).Update("Comment", comment)
// }
