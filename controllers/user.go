package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"challenge/models"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) GetPing(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"value": "pong"})
	return
}

func (u UserController) GetValueByName(c *gin.Context){
	user := c.Params.ByName("name")
	value, ok := userModel.GetByName(user)
	if ok != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": ok})
		c.Abort()
		return
	}
	if value != nil {
		c.JSON(http.StatusOK, gin.H{"user": value.Name, "value": value.Value})
	} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	return
}

func (u UserController) GetUsers(c *gin.Context){
	
	value, ok := userModel.GetAll()
	if ok != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": ok})
		c.Abort()
		return
	} else{
		c.JSON(http.StatusOK, gin.H{"user": "all", "value": value})
	}
		
	return
}