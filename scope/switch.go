package scope

import (
	"gopkg.in/AlecAivazis/survey.v1"
	"fmt"
)

type Switch struct {

}


func Switcher(){
	var name string

	// Ask name
	prompt := &survey.Input{
		Message: "what is your name?",
	}
	survey.AskOne(prompt, &name, nil)

	switch name {
	case "Nik":
		fmt.Println("Hello Nikki")
		break
	case "Jane":
		fmt.Println("Hello Jenni")
		break
	default:
		fmt.Printf("Hello incognito + %v", name)
	}
}
