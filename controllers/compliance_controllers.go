package controllers

import (
	"log"
	"unbound/models"
	"unbound/services"

	"github.com/gin-gonic/gin"
)

func SubmitKYB(c *gin.Context) {

	var kybData models.KYBSubmission
	userID, exists := c.Get("userID")
	log.Println(userID)

	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// veriico se o json esta batendo
	if err := c.ShouldBindJSON(&kybData); err != nil {
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
