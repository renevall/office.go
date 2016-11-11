package auth

import (
	"errors"

	"github.com/ReneVallecillo/office.go/domain"

	"fmt"
)

//AuthService helps with dependency injection and decoupling
type AuthService struct {
	UserRepository domain.UserRepository
}

// type Authorizer interface{
// 	Authorize(r *http.Request)error
// }

//Login logs the user
func (auth *AuthService) Login(email, pass string) (*domain.User, error) {
	fmt.Println("login method reached")
	var authUser *domain.User
	user, err := auth.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if CompareHash(pass, user.Password) {
		authUser, err = auth.UserRepository.FindByID(user.ID)
		if err != nil {
			return nil, err
		}
		authUser.Token = GenerateToken(*authUser)
	} else {
		err = errors.New("hash not equal")
		return nil, err
	}

	return authUser, nil

}