package main

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {
	// -----
	// -----
	sayHello("Mike")
	sayHello("")
	sayHello("Jane")
}

func sayHello(name string)  {
	if name == "" {
		prompt := &survey.Input{
			Message: "what is your name?",
		}
		survey.AskOne(prompt, &name, nil)
	}

	message := "Hello " + name
	fmt.Println(message)
}