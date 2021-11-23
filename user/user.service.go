package user

import (
	"errors"
	"fmt"
	"log"
	"pustaka-api/dto"
	"pustaka-api/models"

	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type IUserService interface {
	Update(user dto.UserUpdateDTO) (models.User, bool, error)
	Profile(userID string) models.User
}

type UserService struct {
	userRepository IUserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (service *UserService) Update(user dto.UserUpdateDTO) (models.User, bool, error) {
	u, isEmailExist, _ := service.userRepository.FindByEmail(user.Email)
	
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	//* jika id  sama, maka email sama tidak apa2
	if u.ID == user.ID {
		updatedUser := service.userRepository.UpdateUser(userToUpdate)
		fmt.Println(updatedUser, "KONTOL")
		return updatedUser, true, nil

	} else {
			//* jika id tidak sama, maka email sama tidak boleh sama, artinya email sudah ada milik org lain
		if isEmailExist {
			return models.User{}, false, errors.New("email already exist")
		}
		updatedUser := service.userRepository.UpdateUser(userToUpdate)
		return updatedUser, true, nil
	}
}

func (service *UserService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}