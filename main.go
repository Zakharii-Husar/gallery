package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.POST("/sign_up", func(c *gin.Context) {
        c.String(200, "sign_up")
    })
    r.GET("/sign_in", func(c *gin.Context) {
        c.String(200, "sign_in")
    })

    r.GET("/users", func(c *gin.Context) {
        c.String(200, "users")
    })
    
    r.GET("/users/search", func(c *gin.Context) {
        query, ok := c.GetQuery("query")
        if ok {
            c.String(200, "You search for " + query)
        }else{
            c.String(200, "No value was provided in search query!")
        }
    })

    r.GET("/users/:user_id", func(c *gin.Context) {
        userID := c.Param("user_id")
        c.String(200, "User ID: " + userID)
    })

    r.GET("/users/:user_id/photos", func(c *gin.Context) {
        c.String(200, "Photos")
    })

    r.GET("/users/:user_id/photos/:photo_id/download", func(c *gin.Context) {
        userID := c.Param("user_id")
        photoID := c.Param("photo_id")
        c.String(200, "User ID: " + userID + " Photo ID: " + photoID + " Download")
    })

    r.POST("/users/:user_id/photos/:photo_id/upload", func(c *gin.Context) {
        userID := c.Param("user_id")
        photoID := c.Param("photo_id")
        c.String(200, "User ID: " + userID + " Photo ID: " + photoID + "Upload")
    })

    r.DELETE("/users/:user_id/photos/:photo_id/delete", func(c *gin.Context) {
        userID := c.Param("user_id")
        photoID := c.Param("photo_id")
        c.String(200, "User ID: " + userID + " Photo ID: " + photoID + " Delete")
    })

    r.Run()
}
