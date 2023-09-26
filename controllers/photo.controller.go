package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/app"
	db "github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/database"
	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/models"
)

func UploadPhoto(ctx *gin.Context) {
    var reqBody app.PhotoRequestBody
    userID, _ := ctx.Get("userID")

    file, err := ctx.FormFile("photo_url")

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing photo file"})
        return
    }

    timestamp := time.Now().UnixNano()
    filePhoto := fmt.Sprintf("%d-%s", timestamp, strings.ReplaceAll(file.Filename, " ", ""))
    filename := "http://localhost:8080/api/v1/photos/" + filePhoto

    reqBody.Title = ctx.PostForm("title")
    reqBody.Caption = ctx.PostForm("caption")
    reqBody.PhotoUrl = filename
    
    photo := models.Photo{
        Title:    reqBody.Title,
        Caption:  reqBody.Caption,
        PhotoUrl: reqBody.PhotoUrl,
        UserID:   userID.(string),
    }

    db := db.Init()

    err = ctx.SaveUploadedFile(file, "./uploads/"+filePhoto)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    if err := db.Create(&photo).Error; err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Photo uploaded successfully", 
        "success": true, 
    })
}

func GetPhoto(ctx *gin.Context) {
    userID, _ := ctx.Get("userID")

    db := db.Init()

    var photos []models.Photo
    result := db.Where("user_id = ?", userID).Preload("User").Find(&photos)

    if result.Error != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Failed to fetch photos",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Successfully retrieve data",
        "data":    photos,
    })
}

func UpdatePhoto(ctx *gin.Context) {
    photoID := ctx.Param("id")
    userID, _ := ctx.Get("userID")

    file, err := ctx.FormFile("photo_url")
    
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing photo file"})
        return
    }

    timestamp := time.Now().UnixNano()
    filePhoto := fmt.Sprintf("%d-%s", timestamp, strings.ReplaceAll(file.Filename, " ", ""))
    filename := "http://localhost:8080/api/v1/photos/" + filePhoto

    updateData := app.PhotoRequestBody{
        Title: ctx.PostForm("title"),
        Caption: ctx.PostForm("caption"),
        PhotoUrl: filename,
    }

    photo := models.Photo{
        Title: updateData.Title,
        Caption: updateData.Caption,
        PhotoUrl: updateData.PhotoUrl,
    }

    db := db.Init()
    result := db.Where("id = ?", photoID).Where("user_id", userID).Preload("User").First(&photo)

    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "message": "Photo not found",
            "success": false,
        })
        return
    }

    if err := db.Save(&photo).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Photo updated successfully",
        "data":    photo,
    })
}

func DeletePhoto(ctx *gin.Context) {
    photoID := ctx.Param("id")
    userID, _ := ctx.Get("userID")

    var photo models.Photo
    
    db := db.Init()

    if err := db.Where("id = ?", photoID).Where("user_id = ?", userID).First(&photo).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "message": "Photo not found",
        })
        return
    }

    if err := db.Where("id = ?", photoID).Where("user_id = ?", userID).Unscoped().Delete(&photo).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Failed to delete photo",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Photo successfully deleted",
    })
}

