package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetUserPhotos(c *gin.Context) {
    skip := c.DefaultQuery("skip", "1")
    take := c.DefaultQuery("take", "20")
    c.String(200, "Photos, skip: " + skip + " take " + take) 
}

func DownloadPhoto(c *gin.Context) {
    userID := c.Param("user_id")
    photoID := c.Param("photo_id")
    c.String(200, "User ID: " + userID + " Photo ID: " + photoID + " Download")
}

func UploadPhoto(c *gin.Context) {
    userID := c.Param("user_id")
    photoID := c.Param("photo_id")
    c.String(200, "User ID: " + userID + " Photo ID: " + photoID + " Upload")
}

func DeletePhoto(c *gin.Context) {
    userID := c.Param("user_id")
    photoID := c.Param("photo_id")
    c.String(200, "User ID: " + userID + " Photo ID: " + photoID + " Delete")
}
