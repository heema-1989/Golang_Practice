package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name  string
	Email string `gorm:"varchar(100);unique_index"`
}
type Book struct {
	Title      string
	Author     string
	CallNUmber int `gorm:"unique_index"`
	PersonId   int `gorm:"primaryKey"`
}

var DB *gorm.DB
var err error

const DNS = "heema:Simform@123@tcp(127.0.0.1:3306)/book_keeper?charset=utf8&parseTime=True&loc=Local"

var (
	people = Person{Name: "Heema Goratela", Email: "heemag2002@gmail.com"}
	books  = []Book{{Title: "Circe", Author: "Madeline Miller", CallNUmber: 1234, PersonId: 1}}
)

func main() {
	//Fetching the env variables
	/*dialect := os.Getenv("DIALECT")
	fmt.Println(dialect)
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	name := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	dbURI := fmt.Sprintf("host:%s, user:%s, dbname:%s, sslmode=disable, password:%s, port:%s", host, user, name, password, dbport)
	fmt.Println(dbURI)
	//connecting to database
	//DB, err = gorm.Open(dialect,dbURI)*/
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	} else {
		fmt.Println("Successfully connected to database")
	}
	DB.AutoMigrate(&Person{})
	DB.AutoMigrate(&Book{})

	temp := []Book{{Title: "The Kite Runner", Author: "Khaled Hosseini", CallNUmber: 3456, PersonId: 2}, {Title: "Flowers of Algernon", Author: "Daniel Keyes", CallNUmber: 6789, PersonId: 3}}
	books = append(books, temp...)
	DB.Create(&people)
	for i, _ := range books {
		DB.Create(&books[i])
	}
	//DB.Delete(&books)

}
