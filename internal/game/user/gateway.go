package user

import (
	"fmt"
	"github.com/tozastation/ca-tech-dojo/pkg/mysql"
)

type (
	Repository interface{
		Create(user User) error
		Get(token string) (*User, error)
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