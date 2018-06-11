package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ORM struct {
	DB	*gorm.DB
}


func NewORMService() (db *ORM, error error){
	orm := new(ORM)
	err := orm.connect()
	if err != nil{
		return nil, err
	}
	return orm, nil
}

func (service *ORM) connect() error{
	host 	 := "localhost"
	port	 := "3306"
	username := "root"
	password := "root"
	dbname   := "student"

	config := string(username + ":" + password + "@tcp(" + host + ":"+ port + ")/" + dbname + "?tls=skip-verify&autocommit=true")
	db, err := gorm.Open("mysql", config)
	if err != nil{
		return err
	}

	service.DB = db

	//defer service.DB.Close()

	return nil
}
