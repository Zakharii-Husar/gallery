package data

import (
	"database/sql"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var DB *sql.DB

func Init() {
    var err error
    connStr := "Server=MSI\\SQLEXPRESS,1433;Database=gallery;Integrated Security=True;"
    DB, err = sql.Open("sqlserver", connStr)
    if err != nil {
        log.Fatal("Unable to connect. Error: ", err)
    }

    createTables()
}

func createTables() {
    _, err := DB.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Users' AND xtype='U')
    CREATE TABLE Users (
        userId INT PRIMARY KEY IDENTITY(1,1),
        username VARCHAR(100) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL
    );`)
    if err != nil {
        log.Fatal(err)
    }

    _, err = DB.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Albums' AND xtype='U')
    CREATE TABLE Albums (
        albumId INT PRIMARY KEY IDENTITY(1,1),
        albumName VARCHAR(255) NOT NULL,
        userId INT,
        FOREIGN KEY (userID) REFERENCES users(id)
    );`)
    if err != nil {
        log.Fatal("Error creating albums table: ", err)
    }

    _, err = DB.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Photos' AND xtype='U')
    CREATE TABLE Photos (
        photoId INT PRIMARY KEY IDENTITY(1,1),
        photoName VARCHAR(255) NOT NULL,
        uploadDate DATETIME NOT NULL,
        albumId INT,
        FOREIGN KEY (albumID) REFERENCES albums(albumID)
    );`)
    if err != nil {
        log.Fatal("Error creating photos table: ", err)
    }
}
