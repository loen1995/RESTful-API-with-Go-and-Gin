package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// game represents data about a videogame
type game struct {
	ID 		string 	`json:"id"`
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
	router.GET("/collection/:id", getGameByID)
	router.POST("/collection", postGame)

	router.Run("localhost:8080")
}

// getCollection responds with the list of all videogames as JSON
func getCollection(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, collection)
}

// postGame adds a game from JSON received in the request body
func postGame(c *gin.Context)  {
	var newGame game

	//Call BindJSON to bind the received JSON to newGame
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	//Add the new game to the slice.
	collection = append(collection, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

// getGameByID locates the game whose ID value matches the id
//parameter sent by the client, then returns that album as a response.
func getGameByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the collection, looking for a game whose ID value
	//matches the parameter.
	for _, a := range collection {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
}