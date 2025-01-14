package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type record struct {
	ID       string `json:"id"` // если с маленькой буквы -- то не экспортируется в json
	PillName string `json:"pillName"`
	DueDate  string `json:"dueDate"` // дата, когда нужно принять препарат
	Status   bool   `json:"status"`
}

// records slice to seed record album data.
var records = []record{
	{ID: "1", PillName: "Yarina", DueDate: "2024-11-28", Status: true},
	{ID: "2", PillName: "Cetrine", DueDate: "2024-11-29", Status: false},
	{ID: "3", PillName: "Berocca", DueDate: "2024-11-27", Status: true},
}

func main() {
	router := gin.Default()
	router.GET("/records", getRecords)
	router.POST("/records", createRecord)
	router.GET("/records/:id", getRecordByID)
	router.DELETE("/records/:id", deleteRecordByID)
	router.PATCH("/records/:id", updateRecordByID)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getRecords(c *gin.Context) {
	// Получение query параметров
	pillName := c.Query("pillName")
	dueDate := c.Query("dueDate")
	statusStr := c.Query("status")
	status, _ := strconv.ParseBool(statusStr)
	filteredRecords := records
	if pillName != "" {
		filteredRecords = filterByName(filteredRecords, pillName)
	}
	if dueDate != "" {
		filteredRecords = filterByDate(filteredRecords, dueDate)
	}
	if statusStr != "" {
		filteredRecords = filterByStatus(filteredRecords, status)
	}
	c.IndentedJSON(http.StatusOK, filteredRecords)
}

// createRecord adds an album from JSON received in the request body.
func createRecord(c *gin.Context) {
	var newRecord record

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newRecord); err != nil {
		return
	}
	// Check if record exists
	for _, a := range records {
		if a.ID == newRecord.ID {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "record already exists"})
			return
		}
	}
	// Add the new album to the slice.
	records = append(records, newRecord)
	c.IndentedJSON(http.StatusCreated, newRecord)
}

// getRecordByID locates the album whose ID value matches the id
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

// deleteRecordByID deletes a record from the list
func deleteRecordByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range records {
		if a.ID == id {
			records = append(records[:i], records[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "record deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}

// updateRecordByID updates a record from the list
func updateRecordByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range records {
		if a.ID == id {
			var updatedRecord record
			if err := c.BindJSON(&updatedRecord); err != nil {
				return
			}
			records[i] = record{
				ID:       id,
				PillName: updatedRecord.PillName,
				DueDate:  updatedRecord.DueDate,
				Status:   updatedRecord.Status,
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "record updated"})
			return
		}
	}
}
