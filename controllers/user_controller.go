package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jcmartins81/webapi-with-go/database"
	"github.com/jcmartins81/webapi-with-go/models"
	"github.com/jcmartins81/webapi-with-go/services"
)

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	p.Password = services.SHA256Encoder(p.Password)

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
