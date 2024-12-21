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

	r.POST("/table/persons/:id", func(c *gin.Context) {
		id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return
		}
		id := uint(id64)

		var newPerson internal.Person

		if err := c.BindJSON(&newPerson); err != nil {
			return
		}

		result, person := internal.UpdatePerson(db, id, newPerson)
		if result.RowsAffected == 0 {
			c.IndentedJSON(http.StatusNotFound, person)
		} else {
			c.IndentedJSON(http.StatusOK, person)
		}
	})

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
