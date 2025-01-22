package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRecords(t *testing.T) {
	// records slice to seed record album data.
	var records = []record{
		{ID: "1", PillName: "Yarina", DueDate: "2024-11-28", Status: true},
		{ID: "2", PillName: "Cetrine", DueDate: "2024-11-29", Status: false},
		{ID: "3", PillName: "Berocca", DueDate: "2024-11-27", Status: true},
	}
	// Создаем тестовый маршрутизатор
	router := gin.Default()
	router.GET("/records", getRecords)

	// Создаем тестовый запрос
	req, err := http.NewRequest("GET", "/records", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем тестовый рекордер для записи ответа
	w := httptest.NewRecorder()

	// Обрабатываем запрос
	router.ServeHTTP(w, req)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, w.Code)
	re := regexp.MustCompile(`\s+`)
	recordsJSON, _ := json.Marshal(records)
	responseString := re.ReplaceAllString(w.Body.String(), "")

	// Проверяем тело ответа
	assert.Equal(t, string(recordsJSON), responseString)
}
