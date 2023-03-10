package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Users struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Books    []Book
}
type Book struct {
	gorm.Model
	Title      string
	Author     string
	IssueDate  time.Time
	ExpiryDate time.Time
	UsersId    int
}

const Dsn = "heema:Simform@123@tcp(127.0.0.1:3306)/library?charset=utf8&parseTime=True&loc=Local"

var tpl *template.Template
var DB *gorm.DB
var err error

func main() {
	users := []Users{
		{Username: "heema-1989", Password: "abc", Email: "heema@gmail.com"}, {Username: "dhatri-2007", Password: "123", Email: "dhatri@gmail.com"}, {Username: "khushi-1904", Password: "xyz@123", Email: "khushi@gmail.com"}}
	fmt.Println(users)
	tpl, _ = template.ParseGlob("/templates/*.html")

	DB, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	/*DB.AutoMigrate(&Users{})
	DB.Create(&users)*/

	http.Handle("/", http.FileServer(http.Dir("./templates/")))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/valid", validateUsers)
	//router.HandleFunc("/imp", notImplemented).Methods("POST")
	http.ListenAndServe(":8080", nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("/templates/login.html")
	tmp.ExecuteTemplate(w, "login.html", nil)
}
func validateUsers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user Users
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("Entered username: ", username, "password: ", password)
	//password = getHash([]byte(password))
	result := DB.Where("username = ? AND password = ?", username, password).First(&user)
	fmt.Println(user)
	fmt.Println("Fetched username: ", user.Username, " password: ", user.Password)
	//fmt.Println("Fetched username: ", result.Username, "Password: ", user.Password)
	if result.Error != nil {
		fmt.Println("Please check username and password")
		tpl.ExecuteTemplate(w, "login.html", "check username and password")
		return
	}
}
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
