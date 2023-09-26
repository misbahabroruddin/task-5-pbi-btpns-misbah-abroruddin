package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username  string    `json:"username" gorm:"unique;not null"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"password" gorm:"not null;size:255;check:(LENGTH(password) >= 6)"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
    passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(passwordHash)

    user.Username = html.EscapeString(strings.TrimSpace(user.Username))

    return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
    if tx.Statement.Changed("Password") {
        passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        user.Password = string(passwordHash)
    }

    user.Username = html.EscapeString(strings.TrimSpace(user.Username))

    return nil
}

func (user *User) ComparePassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

