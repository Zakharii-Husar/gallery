package handlers

import (
	"fmt"
	"gallery/services/usersService"

	"github.com/gin-gonic/gin"
)

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