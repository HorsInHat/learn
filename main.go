package main

import (
	"github.com/HorsInHat/learn/scope"
)

func main() {
	//scp := scope.NewScopeStruct()
	//
	//fmt.Println("------ scp.SET ----------")
	//scp.Set()
	//fmt.Println("------ scp.VIEW ----------")
	//scp.View()
	//
	//fmt.Println("------ fmt SCP ----------")
	//fmt.Println(scp.Config)
	//fmt.Println("------ fmt scope.Public ----------")
	//fmt.Println(scope.Public)

	// Создаем структуру
	inh := scope.NewInherit()
	// Вызываем функцию структуры
	inh.ViewScope()
}
