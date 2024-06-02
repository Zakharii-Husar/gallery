package models

type User struct {
	id       int    `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
}

type Photo struct {
	photo_id int    `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
}
