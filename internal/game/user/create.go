package user

import (
	"fmt"
	"github.com/tozastation/ca-tech-dojo/pkg/jwt"
)

type (
	CreateRequest struct {
		Name string `json:"name"`
	}
	CreateResponse struct {
		Token string `json:"token"`
	}
	CreateUsecase interface {
		Execute(req CreateRequest) (*CreateResponse, error)
	}
	CreateInteractor struct {
		UserRepo Repository
	}
)

func NewCreateInteractor(userRepo Repository) CreateUsecase {
	return &CreateInteractor{UserRepo: userRepo}
}

func (c *CreateInteractor) Execute(req CreateRequest) (*CreateResponse, error) {
	token, err := jwt.Create(req.Name, jwt.SecretPath)
	if err != nil {
		return nil, fmt.Errorf("call jwt create %w", err)
	}
	user := &User{
		Name: req.Name,
		Token: token,
	}
	if err := c.UserRepo.Create(*user); err != nil {
		return nil, err
	}
	return &CreateResponse{Token:token}, nil
}