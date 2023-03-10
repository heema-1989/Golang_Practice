package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name  string `json:"Name"`
	Email string `gorm: "varchar(100);unique_index" json:"Email"`
	Books []Book
}
type Book struct {
	gorm.Model
	Title    string
	Author   string
	Price    int
	PersonId int
}

var (
	People = Person{Name: "Heema Goratela", Email: "heemag2002@gmail.com"}
	Books  = []Book{{Title: "Circe", Author: "Madeline Miller", Price: 1200, PersonId: 1}}
)

func Init() {
	temp := []Book{{Title: "The Kite Runner", Author: "Khaled Hosseini", Price: 3000, PersonId: 1}, {Title: "Flowers of Algernon", Author: "Daniel Keyes", Price: 2500, PersonId: 1}}
	Books = append(Books, temp...)

}
