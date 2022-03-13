package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// game represents data about a videogame
type game struct {
	ID 			string 	`json:"id"`
	Title 		string 	`json:"title"`
	Developer 	string 	`json:"developer"`
	Price 		float64 `json:"price"`
}

// collection slice to seed videogames data
var collection = []game{
	{	ID: 		"1",
		Title:		"Pokemon Sword",
		Developer:	"GameFreak",
		Price:		56.90,
	},
	{	ID: 		"2",
		Title:		"The Legend of Zelda: Breath of the Wild",
		Developer:	"Nintendo",
		Price:		69.99,
	},
	{	ID: 		"3",
		Title:		"Elden Ring",
		Developer:	"FromSoftware",
		Price:		64.98,
	},
}

// set up an association in which getCollection handles requests to the /collections endpoint path
func main()  {
	router := gin.Default()
	router.GET("/collection", getCollection)

	router.Run("localhost:8080")
}

// getCollection responds with the list of all videogames as JSON
func getCollection(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, collection)
}