package main

import (
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

func main() {

}
