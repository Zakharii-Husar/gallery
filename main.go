package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/microsoft/go-mssqldb"
)

//var db *sql.DB

type SignUpInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func main() {
    connStr := "Server=MSI\\SQLEXPRESS,1433;Database=gallery;Integrated Security=True;"
    db, err := sql.Open("sqlserver", connStr)
    if err != nil {
        log.Fatal("Unabele to connect. Error: ", err)
    }
    defer db.Close()

    _, err = db.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='users' AND xtype='U')
    CREATE TABLE users (
        id INT PRIMARY KEY IDENTITY(1,1),
        username VARCHAR(100) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL
    );`)
    if err != nil {
        log.Fatal(err)
    }

        // Create albums table
        _, err = db.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='albums' AND xtype='U')
        CREATE TABLE albums (
            albumID INT PRIMARY KEY IDENTITY(1,1),
            album_name VARCHAR(255) NOT NULL,
            userID INT,
            FOREIGN KEY (userID) REFERENCES users(id)
        );`)
        if err != nil {
            log.Fatal("Error creating albums table: ", err)
        }

        _, err = db.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='photos' AND xtype='U')
        CREATE TABLE photos (
            photoID INT PRIMARY KEY IDENTITY(1,1),
            photo_name VARCHAR(255) NOT NULL,
            upload_date DATETIME NOT NULL,
            albumID INT,
            FOREIGN KEY (albumID) REFERENCES albums(albumID)
        );`)
        if err != nil {
            log.Fatal("Error creating photos table: ", err)
        }

    /////////////////////////////////////////
    r := gin.Default()

    r.POST("/sign_up", func(c *gin.Context) {
        var input SignUpInput
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
                // For demonstration, we'll just return the input as a response
                c.JSON(http.StatusOK, gin.H{
                    "Status": "Success",
                    "data":    input,
                })
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
