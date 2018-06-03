package models

import "github.com/revel/revel"

type UserModel struct {

}

type GuestForm struct {
	Username string
	Password string
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
func IsLogin(ctrl *revel.Controller) bool{
	check := ctrl.Session["isLogin"]

	if check == ""{
		return false
	}
	return true
}