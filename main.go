package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

func main() {
    connStr := "Server=MSI\\SQLEXPRESS,1433;Database=chat;Integrated Security=True;"

    db, err := sql.Open("sqlserver", connStr)

    defer db.Close()
    
    if err != nil {
        log.Fatal("Unabele to connect. Error: ", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    r := gin.Default()

    r.POST("/sign_up", func(c *gin.Context) {
        c.String(200, "sign_up")
    })
    r.GET("/sign_in", func(c *gin.Context) {
        c.String(200, "sign_in")
    })

    r.GET("/users", func(c *gin.Context) {
        skip := c.DefaultQuery("skip", "1")
        take := c.DefaultQuery("take", "20")
        // Convert skip and take to integers
        //pageNum, _ := strconv.Atoi(skip)
        //limitNum, _ := strconv.Atoi(take)
        c.String(200, "Skip: " + skip + "Take: " + take )
    })
    
    r.GET("/users/search", func(c *gin.Context) {
        query, ok := c.GetQuery("query")
        skip := c.DefaultQuery("skip", "1")
        take := c.DefaultQuery("take", "20")
        // Convert skip and take to integers
        //pageNum, _ := strconv.Atoi(skip)
        //limitNum, _ := strconv.Atoi(take)
        if ok {
            c.String(200, "You search for " + query + "Skip " + skip + "Take " + take)
        }else{
            c.String(200, "No value was provided in search query!")
        }
    })

    r.GET("/users/:user_id", func(c *gin.Context) {
        userID := c.Param("user_id")
        c.String(200, "User ID: " + userID)
    })

    r.GET("/users/:user_id/photos", func(c *gin.Context) {
        skip := c.DefaultQuery("skip", "1")
        take := c.DefaultQuery("take", "20")
        // Convert skip and take to integers
        //pageNum, _ := strconv.Atoi(skip)
        //limitNum, _ := strconv.Atoi(take)
        c.String(200, "Photos, skip: " + skip + "take " + take) 
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
