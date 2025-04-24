package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomDate struct {
	time.Time
}

// Запись о таблетке
type record struct {
	ID       string     `json:"id"` // если с маленькой буквы -- то не экспортируется в json
	PillName string     `json:"pillName"`
	DueDate  CustomDate `json:"dueDate"` // дата, когда нужно принять препарат
	Status   bool       `json:"status"`
}

// Реализуем UnmarshalJSON для парсинга строки в дату
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	cd.Time = parsedTime
	return nil
}

var records []record

// Инициализация данных (вызывается автоматически при старте программы)
func init() {
	jsonData := `[
		{"id": "1", "pillName": "Yarina", "dueDate": "2024-11-29", "status": true},
		{"id": "2", "pillName": "Cetrine", "dueDate": "2024-11-28", "status": false},
		{"id": "3", "pillName": "Berocca", "dueDate": "2024-11-30", "status": true}
	]`

	if err := json.Unmarshal([]byte(jsonData), &records); err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}
}

// mustParseDate parses a date string and panics if there is an error.
func mustParseDate(dateStr string) (CustomDate, error) {
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return CustomDate{}, err
	}
	return CustomDate{parsedTime}, nil
}

func main() {
	router := gin.Default()
	router.GET("/records", getRecords)
	router.POST("/records", createRecord)
	router.GET("/records/:id", getRecordByID)
	router.DELETE("/records/:id", deleteRecordByID)
	router.PATCH("/records/:id", updateRecordByID)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// getAlbums responds with the list of all albums as JSON.
func getRecords(c *gin.Context) {
	// Получение query параметров
	pillName := c.Query("pillName")
	dueDate := c.Query("dueDate")
	fromDate := c.Query("fromDate")
	toDate := c.Query("toDate")
	statusStr := c.Query("status")
	// Трансформация входных параметров
	filteredRecords := records
	if pillName != "" {
		filteredRecords = filterByName(filteredRecords, pillName)
	}
	if dueDate != "" {
		date, error := mustParseDate(dueDate)
		if error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": error.Error()})
			return
		}
		filteredRecords = filterByDate(filteredRecords, date)
	}
	if statusStr != "" {
		status, error := strconv.ParseBool(statusStr)
		if error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": error.Error()})
			return
		}
		filteredRecords = filterByStatus(filteredRecords, status)
	}
	if fromDate != "" && toDate != "" {
		from, error := mustParseDate(fromDate)
		to, error := mustParseDate(toDate)
		if error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": error.Error()})
			return
		}
		filteredRecords = filterByInterval(filteredRecords, from, to)
	}
	c.IndentedJSON(http.StatusOK, filteredRecords)
}

// createRecord adds an album from JSON received in the request body.
func createRecord(c *gin.Context) {
	var newRecord record

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newRecord); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
	// c.Request.Body
	// dueDate := c.Body("dueDate")
	// fromDate := c.Query("fromDate")
	// toDate := c.Query("toDate")
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range records {
		if a.ID == id {
			var updatedRecord record
			if err := c.BindJSON(&updatedRecord); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
