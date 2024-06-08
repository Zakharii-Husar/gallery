package usersService

import (
	"gallery/models/DTOs"
	"gallery/models/ORMs"
	"gallery/repos/usersRepo"
)

func Register (input DTOs.SignUpInput){
	user := input.ToORM()
	usersRepo.CreateUser(user)
}

func Login (input DTOs.SignInInput) {
	
}

func GetUserByUname (uname string) *ORMs.User{
	user, err := usersRepo.GetUserByUname(uname)
	if(err == nil){
		return user;
	}
	return nil
}