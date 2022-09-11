package handler

import (
	. "DirectoryApi/dataprocess"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"DirectoryApi/models"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var users []models.User = Loadusers()
	page := models.Page{ID: 1, Name: "Kullanicilar", Description: "Kullanici listesi", URI: "/users"}
	fmt.Println(users)

	userViewModel := models.UserViewModel{Page: page, Users: users}
	t, _ := template.ParseFiles("template/users.html")
	t.Execute(w, userViewModel)

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	no, firstname, lastname, gender := "", "", "", ""
	no = r.FormValue("no")
	No, _ := strconv.Atoi(no)
	firstname = r.FormValue("Firstname")
	lastname = r.FormValue("Lastname")
	gender = r.FormValue("Gender")

	var user models.User = models.User{No: No, Firstname: firstname, Lastname: lastname, Gender: gender}

	Saveuser(user)

	page := models.Page{ID: 2, Name: "Kaydol", Description: "Kayıt Ol", URI: "/users/signup"}
	t, _ := template.ParseFiles("template/signup.html")
	t.Execute(w, page)

}
