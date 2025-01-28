package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reversed-R/time-adjustment-server/internal"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := internal.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}

	internal.InitDB(db)

	r := gin.Default()
	v1 := r.Group("/api/v1")

	// r.GET("/rooms", func(c *gin.Context) {
	// 	c.IndentedJSON(http.StatusOK, internal.GetTable(db))
	// })

	v1.POST("/rooms", func(c *gin.Context) {
		registerRoom(c, db)
	})

	v1.GET("/rooms/:roomId", func(c *gin.Context) {
		getRoom(c, db)
	})

	// v1.GET("/rooms/:roomId/users", func(c *gin.Context) {
	// 	getUsersInRoom(c, db)
	// })

	v1.POST("/rooms/:roomId/users", func(c *gin.Context) {
		registerUser(c, db)
	})

	// v1.POST("/rooms/users/:id", func(c *gin.Context) {
	// 	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	// 	if err != nil {
	// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such uri resource", "uri": "/table/users/" + c.Param("id")})
	// 		return
	// 	}
	// 	id := uint(id64)
	//
	// 	var newUser internal.User
	//
	// 	if err := c.BindJSON(&newUser); err != nil {
	// 		c.IndentedJSON(http.StatusBadRequest, newUser)
	// 		return
	// 	}
	//
	// 	result, user := internal.UpdateUser(db, id, newUser)
	// 	if result.RowsAffected == 0 {
	// 		c.IndentedJSON(http.StatusNotFound, user)
	// 	} else {
	// 		c.IndentedJSON(http.StatusOK, user)
	// 	}
	// })

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}

func registerRoom(c *gin.Context, db *gorm.DB) {
	var newRoomJSON internal.RoomJSON
	var beginTime time.Time
	var newRoom internal.Room

	if err := c.BindJSON(&newRoomJSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newRoomJSON)
		return
	}

	beginTime = time.Date(
		newRoomJSON.BeginTime.Year,
		newRoomJSON.BeginTime.Month,
		newRoomJSON.BeginTime.Day,
		newRoomJSON.BeginTime.Hour,
		newRoomJSON.BeginTime.Min,
		0,
		0,
		time.Local)

	newRoom = internal.Room{
		Name:             newRoomJSON.Name,
		Description:      newRoomJSON.Description,
		BeginTime:        beginTime,
		DayLength:        newRoomJSON.DayLength,
		DayPattern:       newRoomJSON.DayPattern,
		DayPatternLength: newRoomJSON.DayPatternLength,
	}

	result, room := internal.CreateRoom(db, newRoom)
	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusInternalServerError, room)
	} else {
		c.IndentedJSON(http.StatusCreated, room)
	}
}

func getRoom(c *gin.Context, db *gorm.DB) {
	// GET /api/v1/rooms/:roomId
	roomId64, err := strconv.ParseUint(c.Param("roomId"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such uri resource, roomId must be unsigned integer", "uri": "/rooms/" + c.Param("roomId")})
		return
	}
	roomId := uint(roomId64)

	// room, err := internal.GetRoom(db, roomId)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such uri resource, such room Not found", "uri": "/rooms/" + c.Param("roomId")})
	// } else {
	// 	c.IndentedJSON(http.StatusOK, room)
	// }

	roomAllInfo, err := internal.GetRoomAllInfo(db, roomId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error", "uri": "/rooms/" + c.Param("roomId")})
	} else {
		c.IndentedJSON(http.StatusOK, roomAllInfo)
	}
}

// func getUsersInRoom(c *gin.Context, db *gorm.DB) {
// 	roomId64, err := strconv.ParseUint(c.Param("roomId"), 10, 64)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such uri resource, roomId must be unsigned integer", "uri": "/rooms/" + c.Param("roomId")})
// 		return
// 	}
// 	roomId := uint(roomId64)
//
// 	users, err := internal.GetUsersByRoomId(db, roomId)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error", "uri": "/rooms/" + c.Param("roomId") + "/users"})
// 	} else {
// 		c.IndentedJSON(http.StatusOK, users)
// 	}
// }

func registerUser(c *gin.Context, db *gorm.DB) {
	roomId64, err := strconv.ParseUint(c.Param("roomId"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such uri resource, roomId must be unsigned integer", "uri": "/rooms/" + c.Param("roomId")})
		return
	}
	roomId := uint(roomId64)

	type User struct {
		Name    string `json:"name"`
		Comment string `json:"comment"`
	}

	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newUser)
		return
	}

	result, user := internal.CreateUser(db,
		internal.User{
			RoomId:  roomId,
			Name:    newUser.Name,
			Comment: newUser.Comment,
		})

	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusInternalServerError, user)
	} else {
		c.IndentedJSON(http.StatusCreated, user)
	}
}

// func getUsersOfRoom(c *gin.Context, db *gorm.DB) {
// }
