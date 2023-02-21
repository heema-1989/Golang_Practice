package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ID    int `gorm:"primaryKey;"`
	Code  string
	Price uint
}
type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

const DSN = "heema:Simform@123@tcp(127.0.0.1:3306)/crud?charset=utf8&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Successfully connected to database")
	db.AutoMigrate(&Product{})
	//inserting values into tables
	insertProd := &Product{Code: "D42", Price: 4000}
	db.Create(insertProd)
	fmt.Printf("\nINsert ID: %d, Code: %s, Price: %d\n", insertProd.ID, insertProd.Code, insertProd.Price)
	readProd := &Product{}
	db.First(readProd, "code = ?", "D42")
	fmt.Printf("ID: %d, Code: %s, Price: %d\n", readProd.ID, readProd.Code, readProd.Price)
	//BUlk insert query
	var prod []Product = []Product{{Code: "D43", Price: 3000}, {Code: "D44", Price: 5000}, {Code: "D45", Price: 7000}}
	db.Create(&prod)

	prod = []Product{{Code: "D49", Price: 3000}}
	db.Create(&prod)
	for _, val := range prod {
		fmt.Printf("INsert ID: %d, Code: %s, Price: %d\n", val.ID, val.Code, val.Price)
	}
	//Creating users table using struct
	db.AutoMigrate(&User{})
	user := User{Name: "Heema", Age: 20, Birthday: time.Date(2002, 5, 19, 0, 0, 0, 0, time.Local)}
	result := db.Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	//getting first record
	db.First(&prod)
	fmt.Println(&prod)
	db.Take(&prod)
	fmt.Println(&prod)
	db.Last(&prod)
	fmt.Println(&prod)

}
