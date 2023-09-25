package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string 	`json:"username" form:"username"` 
	Password string 	`json:"password" form:"password"`
	Surname  string     `json:"surname"`
	Name     string     `json:"name"`
	Created  time.Time  `json:"datecreated"`
}

func LoadTestUser() *User {
    // Test user with encrypted "test" password.
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 8)
	return &User{
		Username: "test",
		Password: string(hashedPassword), 
		Name: "Test",
		Surname: "User",
		Created: time.Now(),
	}
}