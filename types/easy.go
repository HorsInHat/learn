package types

import (
	"gopkg.in/AlecAivazis/survey.v1"
	"fmt"
)


func SayHello(name string, secondname string)  {
	if name == "" {
		prompt := &survey.Input{
			Message: "what is your name?",
		}
		survey.AskOne(prompt, &name, nil)
	}
	if secondname == ""{
		prompt := &survey.Input{
			Message: "what is your secondname?",
		}
		survey.AskOne(prompt, &secondname, nil)
	}

	message := "Hello " + name + " " + secondname
	fmt.Println(message)
}