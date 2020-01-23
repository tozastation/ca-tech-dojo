package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

// internal variable
var (
	_mysqlDatabase    = os.Getenv("MYSQL_DATABASE")
	_mysqlUser      = os.Getenv("MYSQL_USER")
	_mysqlPassword  = os.Getenv("MYSQL_PASSWORD")
	_mysqlHost      = os.Getenv("MYSQL_HOST")
	_dbParseTime = "parseTime=true"
)

type IMySQLDriver interface {
	GetGorm() *gorm.DB
	OpenMySQLConnectionWithGorm() error
}

type Driver struct {
	Gorm *gorm.DB
}

func NewMySQLDriver() IMySQLDriver {
	return &Driver{}
}

func (d *Driver) OpenMySQLConnectionWithGorm() (err error) {
	connectTemplate := "%s:%s@(%s)/%s?%s"
	connect := fmt.Sprintf(connectTemplate, _mysqlUser, _mysqlPassword, _mysqlHost, _mysqlDatabase, _dbParseTime)
	log.Printf("info: connect is %v", connect)
	d.Gorm, err = gorm.Open("mysql", connect)
	if err != nil {
		return
	}
	return
}

func (d *Driver) GetGorm() *gorm.DB {
	return d.Gorm
}