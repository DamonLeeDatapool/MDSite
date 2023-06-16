package handler

import (
	"fmt"
	"net/http"

	"MServer/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetAllUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
