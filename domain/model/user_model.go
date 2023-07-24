package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type Users []User

func NewUser(name string, age int) (*User, error) {

	user := &User{
		Name: name,
		Age:  age,
	}

	return user, nil
}
