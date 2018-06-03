package controllers

import (
	"github.com/revel/revel"
	"github.com/HorsInHat/learn/web-application/app/models"
)

type User struct {
	Common
}

func (ctrl User) Login(guest models.GuestForm) revel.Result {
	if ctrl.Request.Method == "GET"{
		return ctrl.Render()
	} else {
		var VerifyUser models.GuestForm
		VerifyUser.Username = "admin"
		VerifyUser.Password = "admin"

		guest.Validate(&guest, ctrl.Validation)

		if ctrl.Validation.HasErrors() {
			// Store the validation errors in the flash context and redirect.
			ctrl.Validation.Keep()
			ctrl.FlashParams()
			return ctrl.Redirect(User.Login)
		}

		if guest.Username == VerifyUser.Username{
			if guest.Password == VerifyUser.Password{
				ctrl.Session["isLogin"] = "ok"
				ctrl.Session["username"] = guest.Username
				return ctrl.Redirect("/")
			} else {
				ctrl.Validation.Error("Incorrect Password")
			}
		} else {
			ctrl.Validation.Error("User not found")
		}

		if ctrl.Validation.HasErrors() {
			// Store the validation errors in the flash context and redirect.
			ctrl.Validation.Keep()
			ctrl.FlashParams()
			return ctrl.Redirect(User.Login)
		}

		return ctrl.Render(guest)
	}
}

func (ctrl User) Logout() revel.Result{
	delete(ctrl.Session, "isLogin")
	delete(ctrl.Session, "username")
	return ctrl.Redirect(User.Login)
}