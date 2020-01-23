package user

import (
	"fmt"
	"github.com/tozastation/ca-tech-dojo/pkg/mysql"
	"log"
)

type (
	Repository interface{
		Create(user User) error
		Get(token string) (*User, error)
		Update(user User) error
	}
	Gateway    struct{
		mysql.IMySQLDriver
	}
)

func NewGateway(driver mysql.IMySQLDriver) Repository {
	return &Gateway{IMySQLDriver: driver}
}

func (g *Gateway) Create(user User) error {
	if err := g.IMySQLDriver.GetGorm().Create(&user).Error; err != nil {
		return fmt.Errorf("call gorm create user %w", err)
	}
	return nil
}

func (g *Gateway) Get(token string) (*User, error) {
	user := &User{}
	if err := g.IMySQLDriver.GetGorm().Where("token = ?", token).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (g *Gateway) Update(user User) error {
	log.Printf("debug: update user %v", user)
	if err := g.IMySQLDriver.GetGorm().Save(&user).Error; err != nil {
		return err
	}
	return nil
}