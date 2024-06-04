package handlers

import (
	"fmt"
	"gallery/models/DTOs"
	"gallery/services/usersService"
	"net/http"

	"github.com/gin-gonic/gin"
)


func SignUp(c *gin.Context) {
    var input DTOs.SignUpInput
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    passOk := input.ValidatePass();
    if !passOk{
        c.String(400, "Bad password")
    }
    user := input.ToORM()
    usersService.CreateUser(user);

    c.JSON(http.StatusOK, gin.H{
        "Status": "Success",
        "data":    input,
    })
}

func SignIn(c *gin.Context) {
    c.String(200, "sign_in")
}

func GetUsers(c *gin.Context) {
    skip := c.DefaultQuery("skip", "1")
    take := c.DefaultQuery("take", "20")
    c.String(200, "Skip: " + skip + " Take: " + take)
}

func SearchUsers(c *gin.Context) {
    query, ok := c.GetQuery("query")
    skip := c.DefaultQuery("skip", "1")
    take := c.DefaultQuery("take", "20")

    if ok {
        user := usersService.GetUserByUname(query)
        username := "ZAK"
        fmt.Println(query);
        if user != nil{
            username = user.Username
        }
        c.String(200, "You search for " + username + " Skip " + skip + " Take " + take)
    } else {
        c.String(200, "No value was provided in search query!")
    }
}

func GetUserByID(c *gin.Context) {
    userID := c.Param("user_id")
    c.String(200, "User ID: " + userID)
}

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
