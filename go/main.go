package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reversed-R/time-adjustment-server/internal"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

	r.GET("/table", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, internal.GetTable(db))
	})

	r.POST("/table/users/", func(c *gin.Context) {
		var newUser internal.User

		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, newUser)
			return
		}

		result, user := internal.CreateUser(db, newUser)
		if result.RowsAffected == 0 {
			c.IndentedJSON(http.StatusInternalServerError, user)
		} else {
			c.IndentedJSON(http.StatusCreated, user)
		}
	})

	r.POST("/table/users/:id", func(c *gin.Context) {
		id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such uri resource", "uri": "/table/users/" + c.Param("id")})
			return
		}
		id := uint(id64)

		var newUser internal.User

		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, newUser)
			return
		}

		result, user := internal.UpdateUser(db, id, newUser)
		if result.RowsAffected == 0 {
			c.IndentedJSON(http.StatusNotFound, user)
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	})

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
