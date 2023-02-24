/*package main
type Users struct{
	gorm.Model
	Username string
	Password string
}
var tpl *template.Template
var DB *gorm.DB
func main(){
	users:= []Users{
		{Username: "heema-1989", Password:"abc"}, {Username: "dhatri-2007", Password:"123"}
	}
	
	tpl,_= template.ParseGlobe("templates/*html")
	var err error
	DB,err= gorm.Open(mysql,Open(DNS), &gorm.Config{})
	if err!=nil{
		panic(err)
	}else{
		fmt.Println("Successfully connected to database")
	}
	DB.AutoMigrate(&Users{})
	DB.Create(&users)
	defer DB.Close()

	http.HandleFunc("/login",loginHandler)
	http.ListenAndserve(":3000",nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"login.html",nil)
}
func validateUsers(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var user User{}
	username:= r.FormValue("username")
	password:= r.FormValue("password")
	fmt.println("username: ",username,"password: ",password)
	DB.Where("username = ?, password = ?","username,password").Find(&user)
}
*/