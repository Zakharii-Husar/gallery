package usersService

import (
	"gallery/models/DTOs"
	"gallery/repos/usersRepo"
)

func AddUser (input DTOs.SignUpInput){
	user := input.ToORM()
	usersRepo.AddUser(user)
}