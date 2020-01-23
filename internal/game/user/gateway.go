package user

import (
	"fmt"
	"github.com/tozastation/ca-tech-dojo/pkg/mysql"
)

type (
	Repository interface{
		Create(user User) error
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