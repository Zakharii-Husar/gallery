package usersService

import (
	"gallery/models/ORMs"
	"gallery/repos/usersRepo"
)

func Register (user ORMs.User){
	usersRepo.CreateUser(user)
}

func GetUserByUname (uname string) *ORMs.User{
	user, err := usersRepo.GetUserByUname(uname)
	if(err == nil){
		return user;
	}
	return nil
}