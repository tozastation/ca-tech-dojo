package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IController interface {
		Create(c *gin.Context)
		Get(c *gin.Context)
	}
	Controller struct {
		UserCreate CreateUsecase
		UserGet GetUsecase
	}
)

func NewUserController(userCreate CreateUsecase, userGet GetUsecase) IController {
	return &Controller{UserCreate:userCreate, UserGet:userGet}
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

func (c2 Controller) Get(c *gin.Context) {
	token := c.GetHeader("x-token")
	if !IsExistString(token) {
		c.String(http.StatusBadRequest, "please set x-token value")
		return
	}
	resp, err := c2.UserGet.Execute(token)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func IsExistString(a string) bool {
	if a == "" {
		return false
	}
	return true
}