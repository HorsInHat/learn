package scope

import "fmt"

type Inherit struct {
	Scope Scope
}

func NewInherit() *Inherit{
	inh := new(Inherit)
	return inh
}

func (self *Inherit) ViewScope()  {
	self.Scope.Set()
	// Присваиваем возраст ребенку ребенка
	//self.Scope.Child.Age = 8
	fmt.Println(self)
}