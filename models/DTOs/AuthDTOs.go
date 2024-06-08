package DTOs

import (
	"gallery/models/ORMs"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type SignUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
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

func (input *SignUpInput) ValidatePass() bool {
	pass := input.Password

	isLong := len(pass) > 5
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(pass)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pass)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(pass)

	if !isLong || !hasLower || !hasUpper || !hasNumber{
		return false
	}
    return true
}

func (input *SignInInput) VerifyPass(hashedPass string) bool {

hashedPasswordBytes := []byte(hashedPass)
plainPasswordBytes := []byte(input.Password)

err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, plainPasswordBytes)
    if err != nil {
		return false
    }
	return true
}
