package DTOs

import (
	"gallery/models/ORMs"

	"golang.org/x/crypto/bcrypt"
)

type SignUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (input *SignUpInput) ToORM() ORMs.User {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    return ORMs.User{
        Username: input.Username,
        Password: string(hash),
    }
}
