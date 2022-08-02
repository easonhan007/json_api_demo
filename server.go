package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type authHeader struct {
	Token string `header:"token"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var players = []player{
	{ID: "1", Name: "Ronaldo Cristiano", Age: 37},
	{ID: "2", Name: "De Gea David", Age: 31},
	{ID: "3", Name: "Eriksen Christian", Age: 30},
	{ID: "4", Name: "Rashford Marcus", Age: 24},
	{ID: "5", Name: "Maguire Harry", Age: 29},
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/api/v1/players", getPlayers)
	router.GET("/api/v2/players", getPlayersV2)
	router.GET("/api/v1/albums", getAlbums)
	router.GET("/api/v1/albums/:id", getAlbumByID)
	router.POST("/api/v1/albums", postAlbums)
	router.DELETE("/api/v1/albums/:id", removeAlbumByID)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "pong",
	})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getPlayers(c *gin.Context) {
	token := c.DefaultQuery("token", "NOT FOUND")
	if strings.ToLower(token) == "secert" {
		c.IndentedJSON(http.StatusOK, players)
		return
	}

	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "access token required"})
}

func getPlayersV2(c *gin.Context) {
	h := authHeader{}
	err_msg := "access token required"

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err_msg})
	}
	if strings.ToLower(h.Token) == "secert" {
		c.IndentedJSON(http.StatusOK, players)
		return
	}

	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err_msg})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func removeAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for index, a := range albums {
		if a.ID == id {
			tmp := albums[index]
			albums[index] = albums[len(albums)-1]
			albums = albums[:len(albums)-1]
			c.IndentedJSON(http.StatusOK, tmp)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
