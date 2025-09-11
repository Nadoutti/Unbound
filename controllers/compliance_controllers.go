package controllers

import (
	"github.com/gin-gonic/gin"
	"unbound/models"
	"unbound/services"
)

func SubmitKYB(c *gin.Context) {

	var kybData models.KYBSubmission
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// veriico se o json esta batendo
	if err := c.ShouldBindBodyWithJSON(&kybData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// chama o service de registro
	data, err := services.SubmitKYB(&kybData, userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, data)

}
