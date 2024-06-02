package models

type User struct {
	userId   int
	username string
	password string
}

type Album struct {
	albumId   int
	akbumName string
	userId    string
}

type Photo struct {
	photoId    int
	photoName  string
	uploadDate string
	albumId    int
}
