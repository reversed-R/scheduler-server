package internal

import (
	"gorm.io/gorm"
	"time"
)

// length of days : all has length of days, begin day
// time pattern in a day that should be supported : all has length of times
// a day (1)
// morning and afternoon (2)
// classes time of tsukuba univ (1 + 6 + 1)
// an hour cycle, begin and end can be freely decided (need to have begin time) (variable)

// day pattern:
//   "A_DAY"
//   "AM_AND_PM"
//   "CLASSES_OF_TSUKUBA_UNIV"
//   "HOURS"

const (
	ADAY                    = "ADAY"
	AM_AND_PM               = "AM_AND_PM"
	CLASSES_OF_TSUKUBA_UNIV = "CLASSES_OF_TSUKUBA_UNIV"
	HOURS                   = "HOURS"
)

const (
	OK      = "OK"
	NOT_BAD = "NOT_BAD"
	BAD     = "BAD"
)

func DayPatternToLength(dayPattern string) int {
	patternMap := map[string]int{
		ADAY:                    1,
		AM_AND_PM:               2,
		CLASSES_OF_TSUKUBA_UNIV: 8,
	}

	return patternMap[dayPattern]
}

// for front JSON
type RoomAllInfoJSON struct {
	RoomName         string     `json:"roomName"`
	RoomDescription  string     `json:"roomDescription"`
	BeginTime        TimeJSON   `json:"beginTime"`
	DayLength        uint       `json:"dayLength"`
	DayPattern       string     `json:"dayPatternId"`
	DayPatternLength uint       `json:"dayPatternLength"`
	Users            []UserJSON `json:"users"`
}

type UserJSON struct {
	Name           string   `json:"name"`
	Comment        string   `json:"comment"`
	Availabilities []string `json:"availabilities"`
}

type RoomJSON struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	BeginTime        TimeJSON `json:"beginTime"`
	DayLength        uint     `json:"dayLength"`
	DayPattern       string   `json:"dayPatternId"`
	DayPatternLength uint     `json:"dayPatternLength"`
}

type TimeJSON struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
	Hour  int        `json:"hour"`
	Min   int        `json:"min"`
}

// for DB tables

type Room struct {
	gorm.Model
	Name             string    `gorm:"name"`
	Description      string    `gorm:"description"`
	BeginTime        time.Time `gorm:"begin_time"`
	DayLength        uint      `gorm:"day_length"`
	DayPattern       string    `gorm:"day_pattern"`
	DayPatternLength uint      `gorm:"day_pattern_length"`
	// Password   string
}

type User struct {
	gorm.Model
	RoomId  uint   `gorm:"room_id"`
	Name    string `gorm:"name"`
	Comment string `gorm:"comment"`
}

type Plan struct {
	gorm.Model
	UserId       uint   `gorm:"user_id"`
	TimeId       uint   `gorm:"time_id"`
	Availability string `gorm:"availability"`
}
