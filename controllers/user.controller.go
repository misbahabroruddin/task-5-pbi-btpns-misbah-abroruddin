package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/app"
	db "github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/database"
	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/models"
)

func UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	db := db.Init()

	var foundUser models.User
	if err := db.First(&foundUser, userID).Error; err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
		})
		return
	}

	var updatedUser app.UpdateUserInput
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"message": "Bad request",
			})
			return
	}

	foundUser.Username = updatedUser.Username
	foundUser.Email = updatedUser.Email
	foundUser.Password = updatedUser.Password

	if err := db.Save(&foundUser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User updated successfully",
	})
}

func DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	db := db.Init()

	var foundUser models.User
	if err := db.First(&foundUser, userID).Error; err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
		})
		return
	}

	if err := db.Where("id = ?", userID).Unscoped().Delete(&foundUser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to delete user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User successfully deleted",
	})
}