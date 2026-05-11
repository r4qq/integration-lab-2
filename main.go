package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// globalna struktura "zależności" - wskaźnik do definicji bazy danych. Dzięki niej zamiast
// używać globalnej zmiennej DB, każda funkcja będzie metodą tej struktury
type Env struct {
	db *gorm.DB
}

// struktura będąca jednocześnie modelem posta dla db
type Post struct {
	gorm.Model        // dodaje automatycznie pola id, createdAt, updatedAt, deletedAt
	Title      string `gorm:"type:text;not null" form:"title"`
	Body       string `gorm:"type:text;not null" form:"body"`
	UserID     uint
	Author     User `gorm:"foreignKey:UserID"`
}

// struktura będąca jednocześnie modelem użytkownika dla db
type User struct {
	gorm.Model        // dodaje automatycznie pola id, createdAt, updatedAt, deletedAt
	Username   string `gorm:"unique;not null" form:"user"`
	Posts      []Post
}

// helper do generowania połączenia z bazą
func setupDatabase() *gorm.DB {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("connection to db failed: %v", err)
	}

	err = db.AutoMigrate(&User{}, &Post{})
	if err != nil {
		log.Fatalf("migartion to db failed: %v", err)
	}

	// przykladowi userzy
	var users = []User{
		{Username: "glenda"},
		{Username: "pike"},
	}

	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {
		db.Create(&users)
		log.Println("users seeded")
	}

	// przykładowe posty
	var posts = []Post{
		{Title: "9front", Body: "the front fell off", UserID: users[0].ID},
		{Title: "go", Body: "lang", UserID: users[1].ID},
	}

	db.Model(&Post{}).Count(&count)
	if count == 0 {
		db.Create(&posts)
		log.Println("users seeded")
	}
	return db
}

func main() {

}
