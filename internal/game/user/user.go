package user

import "github.com/tozastation/ca-tech-dojo/pkg/gorm"

type User struct {
	gorm.BaseModel
	Name  string `json:"name"`
	Token string `json:"token"`
}
