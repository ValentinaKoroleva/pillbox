package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type record struct {
	ID       string `json:"id"`
	PillName string `json:"pillName"`
	DueDate  string `json:"dueDate"` // дата, когда нужно принять препарат
	Status   bool   `json:"status"`
}

// records slice to seed record album data.
var records = []record{
	{ID: "1", PillName: "Blue Train", DueDate: "2024-11-28", Status: true},
	{ID: "2", PillName: "Jeru", DueDate: "2024-11-29", Status: false},
	{ID: "3", PillName: "Sarah Vaughan and Clifford Brown", DueDate: "2024-11-27", Status: true},
}

func main() {
	router := gin.Default()
	router.GET("/records", getRecords)
	router.POST("/records", createRecord)
	router.GET("/records/:id", getRecordByID)
	// router.DELETE("/records/:id", getRecordByID)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getRecords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, records)
}

// postAlbums adds an album from JSON received in the request body.
func createRecord(c *gin.Context) {
	var newRecord record

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newRecord); err != nil {
		return
	}

	// Add the new album to the slice.
	records = append(records, newRecord)
	c.IndentedJSON(http.StatusCreated, newRecord)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getRecordByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range records {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}
