package main

import (
	"gallery/data"
	"gallery/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    data.Init()
    defer data.DB.Close()

    r := gin.Default()

    r.POST("/sign_up", handlers.SignUp)
    r.GET("/sign_in", handlers.SignIn)
    r.GET("/users", handlers.GetUsers)
    r.GET("/users/search", handlers.SearchUsers)
    r.GET("/users/:user_id", handlers.GetUserByID)
    r.GET("/users/:user_id/photos", handlers.GetUserPhotos)
    r.GET("/users/:user_id/photos/:photo_id/download", handlers.DownloadPhoto)
    r.POST("/users/:user_id/photos/:photo_id/upload", handlers.UploadPhoto)
    r.DELETE("/users/:user_id/photos/:photo_id/delete", handlers.DeletePhoto)

    r.Run()
}
