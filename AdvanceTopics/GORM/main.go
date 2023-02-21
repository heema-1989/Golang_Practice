package main

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	Name  string
	Email string
}

// embedding this author struct in blog struct using embedded tag
type Blog struct {
	ID      int
	Author  Author `gorm:embedded`
	Upvotes int32
}

// gorm.Model definition
type Users struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {

}
