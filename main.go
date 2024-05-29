package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

func main() {
    
    // Context
    ctx, stop := context.WithCancel(context.Background())
    defer stop()

    appSignal := make(chan os.Signal, 3)
    signal.Notify(appSignal, os.Interrupt)
    go func() {
        <-appSignal
        stop()
    }()

    // Replace with your own connection parameters
    var server = "localhost"
    var port = 1433
    var user = "chat_user"
    var password = "ChatUsrPswd_781"
    var database = "chat"

    dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        server, user, password, port, database)

    db, err := sql.Open("sqlserver", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    db.SetConnMaxLifetime(0)
    db.SetMaxIdleConns(3)
    db.SetMaxOpenConns(3)

    // Open connection
    OpenDbConnection(ctx, db)
}

func OpenDbConnection(ctx context.Context, db *sql.DB) {
    ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
    defer cancel()

    if err := db.PingContext(ctx); err != nil {
        log.Fatal("Unable to connect to database: %v", err)
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
