package models

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	"encoding/base64"
	"github.com/HorsInHat/learn/web-application/app/services"
)

type User struct {
	gorm.Model
	Login 		string
	Password 	string
	Name 		string
	SecondName 	string
}

type GuestForm struct {
	Username string
	Password string
}

type UserForm struct {
	Login 	string
	Name 		string
	SecondName 	string
	Password 	string
}

// -------------- FORMS --------------
func (f *GuestForm) Validate(form *GuestForm, v *revel.Validation) {
	v.Required(form.Username).Message("Field Username can't be empty.")
	v.MinSize(form.Username, 3).Message("Username must have at list min size 3 symbols")
	v.MaxSize(form.Username, 25).Message("Username must have at list max size 25 symbols")
	v.Required(form.Password).Message("Field Password can't be empty.")
	v.MinSize(form.Password, 3).Message("Password must have at list min size 3 symbols")
	v.MaxSize(form.Password, 25).Message("Password must have at list max size 25 symbols")
}

// -------------- MODEL --------------
func (model *User) Create(form *UserForm) {
	newUser := User{
		Name: form.Name,
		SecondName: form.SecondName,
		Login: form.Login,
		Password: createPassword(form.Login, form.Password),
	}

	orm, err := services.NewORMService()
	if err != nil{
		panic("Can't create user")
	}
	orm.DB.Create(&newUser)
}

func createPassword(login string, password string) string {
	pass := string(login) + string(password) + string(login)
	hash := base64.StdEncoding.EncodeToString([]byte(pass))
	return hash
}

func IsLogin(ctrl *revel.Controller) bool{
	check := ctrl.Session["isLogin"]

	if check == ""{
		return false
	}
	return true
}
