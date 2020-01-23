package main

import (
	"github.com/tozastation/ca-tech-dojo/internal/game/server"
	"github.com/tozastation/ca-tech-dojo/internal/game/user"
	"github.com/tozastation/ca-tech-dojo/pkg/mysql"
	"log"
)

func main() {
	//DI
	mysqlDriver := mysql.NewMySQLDriver()
	if err := mysqlDriver.OpenMySQLConnectionWithGorm(); err != nil {
		log.Fatalln(err)
	}
	if err := mysqlDriver.GetGorm().AutoMigrate(&user.User{}).Error; err != nil {
		log.Fatalln(err)
	}
	log.Printf("info: create user table\n")
	userGateway := user.NewGateway(mysqlDriver)
	userCreateInteractor := user.NewCreateInteractor(userGateway)
	userGetInteractor := user.NewGetInteractor(userGateway)
	userUpdateInteractor := user.NewUpdateInteractor(userGateway)
	userCtrl := user.NewUserController(userCreateInteractor, userGetInteractor, userUpdateInteractor)
	s := server.NewServer(userCtrl)
	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}
