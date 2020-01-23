package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IController interface {
		Create(c *gin.Context)
	}
	Controller struct {
		UserCreate CreateUsecase
	}
)

func NewUserController(userCreate CreateUsecase) IController {
	return &Controller{UserCreate:userCreate}
}

func (c2 Controller) Create(c *gin.Context) {
	body := &CreateRequest{}
	if err := c.BindJSON(body); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := c2.UserCreate.Execute(*body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, resp)
	return
}