package scope

import "fmt"

var private string
var Public string

type Scope struct {
	Config string
	Child Child
}

type Child struct {
	Name string
	Age int
}

func NewScopeStruct() *Scope{
	scope := new(Scope)
	return scope
}

func (self *Scope) Set() {
	//private = "not empty"
	//Public = "not empty"
	//self.Child.Name = "Nik"
	//
	//self.Config = "not empty"
}

func (self *Scope) View() {
	//self.Config = "empty"
	fmt.Println(self.Config)
}

// TODO(ichiro18): 1) Сделать функции заполнения структуры значениями по-умолчанию для каждой из структур
// TODO(ichiro18): 2) При инициализации структур вызывать эти функции
func (self *Child) SetDefault(){

}