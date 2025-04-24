package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRecords(t *testing.T) {

	var records []record
	jsonData := `[
		{"id": "1", "pillName": "Yarina", "dueDate": "2024-11-29", "status": true},
		{"id": "2", "pillName": "Cetrine", "dueDate": "2024-11-28", "status": false},
		{"id": "3", "pillName": "Berocca", "dueDate": "2024-11-30", "status": true}
	]`

	if err := json.Unmarshal([]byte(jsonData), &records); err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
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

// Benchmark for getRecords
func BenchmarkGetRecords(b *testing.B) {
	// Создаем тестовый маршрутизатор
	router := gin.Default()
	router.GET("/records", getRecords)

	// Создаем тестовый запрос
	req, err := http.NewRequest("GET", "/records", nil)
	if err != nil {
		b.Fatal(err)
	}

	// Создаем тестовый рекордер для записи ответа
	w := httptest.NewRecorder()

	// Обрабатываем запрос
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(w, req)
	}
}
