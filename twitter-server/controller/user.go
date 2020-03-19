package controller

import (
	"cekirdek/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Login(c *gin.Context) {
	var usermodel model.LoginModel
	if err := c.ShouldBind(&usermodel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.Error{
			Message: "Json binding error",
		})
		return
	}

	if err := validate.Struct(&usermodel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.Error{
			Message: "validation error",
		})
		return
	}
}
