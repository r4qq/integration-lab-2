package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func testCreateUser(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{}) //"plik bazy danych tworzy się w pamieci ram"
	db.AutoMigrate(&User{}, &Post{})

	e := Env{db: db}

	gin.SetMode(gin.DebugMode)

	//testowy router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.templ")
	router.POST("/users/new", e.createUser)

	//dane do testowego "formularza"
	formData := url.Values{}
	formData.Set("user", "test_user")
	body := strings.NewReader(formData.Encode())

	responseRecoder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/new", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(responseRecoder, req)

	//test 1: czy po dodaniu przekierwowano na strone główna
	if responseRecoder.Code != http.StatusFound {
		t.Errorf("Oczekiwano kodu 302, otrzymano %d", responseRecoder.Code)
	}

	//test 2: czy rekord jest w bazie
	var user User
	if err := e.db.First(&user, "username = ?", "test_user").Error; err != nil {
		t.Errorf("Nie znaleziono uzytkownika w bazie %s", err.Error())
	}

}
