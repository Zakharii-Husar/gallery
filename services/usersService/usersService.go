package usersService

import (
	"gallery/models/DTOs"
	"gallery/models/ORMs"
	"gallery/repos/usersRepo"
)

func AddUser (input DTOs.SignUpInput){
	user := ORMs.User{
        Username: input.Username,
        Password: input.Password,
    }
	usersRepo.AddUser(user)
}