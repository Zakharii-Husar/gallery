package handlers

import (
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
    unameExists := usersService.GetUserByUname(input.Username)

    if unameExists != nil{
        c.String(400, "Username already registered!")
    }
    user := input.ToORM()
    usersService.Register(user);

    c.JSON(http.StatusOK, gin.H{
        "Status": "Success",
        "data":    input,
    })
}

func SignIn(c *gin.Context) {
	var input DTOs.SignUpInput
    if err := c.BindJSON(&input); err != nil {
        c.JSON(400, gin.H{"Bad request": err.Error()})
        return
    }
    c.String(200, "sign_in")
}