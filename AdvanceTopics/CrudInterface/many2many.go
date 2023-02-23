package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Customer struct {
	gorm.Model

	Orders       []Order `gorm:"many2many:customer_orders"`
	CustomerName string
	Email        string `gorm:"column: email"`
}
type Order struct {
	gorm.Model
	Customers   []Customer `gorm:"many2many:customer_orders"`
	CustomerId  int
	ProductName string
}

var (
	customer = []Customer{{CustomerName: "Heema Goratela", Email: "heemag2002@gmail.com"}, {CustomerName: "Khushi Rajpal", Email: "sbv@gmail.com"}}
	order    = []Order{{ProductName: "pen", CustomerId: 1}, {ProductName: "pencil", CustomerId: 2}, {ProductName: "Crayons", CustomerId: 1}}
)

const Dsn = "heema:Simform@123@tcp(127.0.0.1:3306)/many2many?charset=utf8&parseTime=True&loc=Local"

var Db *gorm.DB
var e error

func main() {
	Db, e = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if e != nil {
		panic("Failed to connect to database")
	} else {
		fmt.Println("Successfully connected to database")
	}
	/*Db.AutoMigrate(&Customer{})
	Db.AutoMigrate(&Order{})
	for i, _ := range customer {

		Db.Create(&customer[i])
	}
	for i, _ := range order {

		Db.Create(&order[i])
	}*/
	router := mux.NewRouter()
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customerOrders", getCustomerOrders).Methods("GET")
	router.HandleFunc("/customerOrder/{id}", getCustomerOrder).Methods("GET")
	router.HandleFunc("create/c", createCustomer).Methods("POST")
	router.HandleFunc("create/o", createOrder).Methods("POST")
	http.ListenAndServe(":8000", router)
}
func getCustomers(w http.ResponseWriter, r *http.Request) {
	var customers = []Customer{}
	Db.Find(&customers)
	json.NewEncoder(w).Encode(&customers)
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	var customer = Customer{}
	params := mux.Vars(r)
	Db.Find(&customer, params["id"])
	json.NewEncoder(w).Encode(&customer)
}
func getCustomerOrders(w http.ResponseWriter, r *http.Request) {

	var customers []Customer

	Db.Model(&customers).Preload("Orders").Find(&customers)
	json.NewEncoder(w).Encode(&customers)
}
func getCustomerOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer = Customer{}
	Db.Model(&customer).Preload("Orders").Find(&customer, params["id"])
	json.NewEncoder(w).Encode(&customer)
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)
	p := DB.Create(&customer)
	e = p.Error
	if e != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&customer)
	}
}
func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Person
	json.NewDecoder(r.Body).Decode(&order)
	p := Db.Create(&order)
	e = p.Error
	if e != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&order)
	}
}
