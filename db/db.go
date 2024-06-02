package db

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
    _, err := DB.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='users' AND xtype='U')
    CREATE TABLE users (
        id INT PRIMARY KEY IDENTITY(1,1),
        username VARCHAR(100) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL
    );`)
    if err != nil {
        log.Fatal(err)
    }

    _, err = DB.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='albums' AND xtype='U')
    CREATE TABLE albums (
        albumID INT PRIMARY KEY IDENTITY(1,1),
        album_name VARCHAR(255) NOT NULL,
        userID INT,
        FOREIGN KEY (userID) REFERENCES users(id)
    );`)
    if err != nil {
        log.Fatal("Error creating albums table: ", err)
    }

    _, err = DB.Exec(`IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='photos' AND xtype='U')
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
}
