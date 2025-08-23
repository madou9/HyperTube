package main

import (
	"log"
	"net/http"
	"os"
	"github.com/madou9/HyperTube/internal/database"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID			int `json:"id" db:"id"`
	FirstName 	string  `json:"first_name" db:"first_name"`
	LastName 	string  `json:"last_name" db:"last_name"`
	Email 		string  `json:"email" db:"email"`
	Password 	string  `json:"password,omitempty" db:"password_hash"`
	CreatedAt 	time.time `json:"created_at" db:"created_at"`
}

type RegisterRequest struct {
	FirstName 	string  `json:"first_name" db:"first_name"`
	LastName 	string  `json:"last_name" db:"last_name"`
	Email 		string  `json:"email" db:"email"` 
	Password 	string  `json:"password,omitempty" db:"password_hash"`
}

type LoginRequest struct {
	Email 		string  `json:"email" db:"email"`
	Password 	string  `json:"password,omitempty" db:"password_hash"`
}
type AuthResponse struct {
	ID			int `json:"id" db:"id"`
	FirstName 	string  `json:"first_name" db:"first_name"`
	LastName 	string  `json:"last_name" db:"last_name"`
	Email 		string  `json:"email" db:"email"`
	Email 		string  `json:"password,omitempty" db:"password_hash"`
	CreatedAt 	time.time `json:"created_at" db:"created_at"`
}

func main() {
	db.InitDb()
	router := gin.Default()

	router.Get("/users", getAllUser)
	router.run()
}

func(c *context) getAllUser() {

}