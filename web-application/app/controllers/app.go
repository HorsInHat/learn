package controllers

import (
	"github.com/revel/revel"
	"github.com/HorsInHat/learn/web-application/app/models"
)

type App struct {
	Common
}

func (ctrl App) Index() revel.Result {
	if models.IsLogin(ctrl.Controller){
		username := ctrl.Session["username"]
		ctrl.ViewArgs["username"] = username

		ctrl.ViewArgs["debug"] = "0"
		return ctrl.Render()
	}
	// TODO(ichiro18): Сделать шаблон для внутренних страниц
	// TODO(ichiro18): Пример: https://getbootstrap.com/docs/4.1/examples/dashboard/
	return ctrl.Redirect(UserController.Login)
}
