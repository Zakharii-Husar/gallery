package usersService

import (
	"gallery/models/ORMs"
	"gallery/repos/usersRepo"
)

func CreateUser (user ORMs.User){
	usersRepo.CreateUser(user)
}

func GetUserByUname (user ORMs.User){
	usersRepo.CreateUser(user)
}