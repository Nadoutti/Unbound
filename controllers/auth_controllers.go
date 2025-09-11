package controllers

import (
	"github.com/gin-gonic/gin"
	"unbound/models"
	"unbound/services"
)

func Login(c *gin.Context) {

	var loginData models.Login

	// veriico se o json esta batendo
	if err := c.ShouldBindBodyWithJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// chama o service de login
	data, err := services.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)

}

// O registro para business e customer sao iguais porque eles colocam os mesmos dados iniciais

func Register(c *gin.Context) {

	var registerData models.Register

	// veriico se o json esta batendo
	if err := c.ShouldBindBodyWithJSON(&registerData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// chama o service de registro
	data, err := services.RegisterUser(registerData.Email, registerData.Password, registerData.Nome, registerData.Phone, registerData.UserType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, data)
}
