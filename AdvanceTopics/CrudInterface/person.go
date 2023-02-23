package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Person struct {
	gorm.Model
	Name  string
	Email string `gorm:"varchar(100);unique_index"`
	Books []Book
}
type Book struct {
	Title      string
	Author     string
	CallNUmber int `gorm:"unique_index"`
	PersonId   int //here it will automatically treat this as a foreign key--classname and Id
}

// these structs have has 2 many relations.
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
	/*DB.AutoMigrate(&Person{})
	DB.AutoMigrate(&Book{})

	temp := []Book{{Title: "The Kite Runner", Author: "Khaled Hosseini", CallNUmber: 3456, PersonId: 1}, {Title: "Flowers of Algernon", Author: "Daniel Keyes", CallNUmber: 6789, PersonId: 1}}
	books = append(books, temp...)
	DB.Create(&people)
	for i, _ := range books {
		DB.Create(&books[i])
	}*/
	//API routes
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/create/person", CreatePerson).Methods("POST")
	router.HandleFunc("/delete/person/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/create/books", CreateBooks).Methods("POST")
	http.ListenAndServe(":8080", router)
}
func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people = []Person{}
	DB.Find(&people)
	//obj, _ := json.MarshalIndent(people, "", "\t")
	//fmt.Println(obj)
	json.NewEncoder(w).Encode(&people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	//var books []Book
	params := mux.Vars(r)
	//getting all the books of a particular person
	//DB.Find(&person, params["id"])
	DB.Model(&person).Preload("Books").Find(&person, params["id"])

	//person.Books = books
	json.NewEncoder(w).Encode(&person)

}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	p := DB.Create(&person)
	err = p.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&person)
	}
}
func CreateBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book

	json.NewDecoder(r.Body).Decode(&books)
	b := DB.Create(&books)
	err = b.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&books)
	}
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	DB.Find(&person, params["id"])
	DB.Delete(&person)
	json.NewEncoder(w).Encode(&person)
}
