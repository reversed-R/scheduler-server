package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reversed-R/time-adjustment-server/internal"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// albums slice to seed record album data.
// var albums = []internal.Album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	db, err := internal.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}

	internal.InitDB(db)

	r := gin.Default()

	// r.GET("/albums/:id", getAlbumById)
	// r.POST("/albums", postAlbums)

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

// func getTable(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, internal.GetTable())
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum internal.Album
//
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
//
// 	albums = append(albums, newAlbum)
// 	internal.Create(newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
//
// func getAlbumById(c *gin.Context) {
// 	id := c.Param("id")
//
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }
