package ORMs

type User struct {
	UserId   int
	Username string
	Password string
}

type Album struct {
	AlbumId   int
	AlbumName string
	userId    string
}

type Photo struct {
	PhotoId    int
	PhotoName  string
	UploadDate string
	AlbumId    int
}