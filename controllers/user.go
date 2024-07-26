package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/models"
)

// GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []models.User

	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		fmt.Println("failed to bind the json", err.Error())
		return
	}

	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	var user models.User

	id := c.Params.ByName("id")

	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user models.User

	id := c.Params.ByName("id")

	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}

	if err := c.BindJSON(&user); err != nil {
		fmt.Println("failed to bind the json", err.Error())
		return
	}

	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user models.User

	id := c.Params.ByName("id")

	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
