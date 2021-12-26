package auth

import (
	"errors"
	"log"
	"pustaka-api/dto"
	"pustaka-api/models"
	"pustaka-api/user"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type IAuthService interface {
	VerifyCredential(email string, password string) (models.User, error)
	CreateUser(user dto.RegisterDTO) models.User
	FindByEmail(email string) (models.User, error)
	IsDuplicateEmail(email string) bool
}

type AuthService struct {
	userRepository user.IUserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep user.IUserRepository) IAuthService {
	return &AuthService{
		userRepository: userRep,
	}
}

func (s *AuthService) VerifyCredential(email string, password string) (models.User, error) {

	user, isEmailExist, errEmailNotFound := s.userRepository.FindByEmail(email)

	isSamePassword, errHash := comparePassword(user.Password, []byte(password))

	if !isEmailExist && !isSamePassword {
		return user, errors.New("invalid credentials")
	}

	if !isEmailExist {
		return user, errEmailNotFound
	}

	if !isSamePassword{
		return user, errHash
	}

	return user, nil
}

func (s *AuthService) CreateUser(user dto.RegisterDTO) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := s.userRepository.InsertUser(userToCreate)
	return res
}

func (s *AuthService) FindByEmail(email string) (models.User, error) {
	user, isEmailExist, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return user, err
	}
	if isEmailExist {
		return user, errors.New("email already exist")
	}
	return user, nil
}

func (s *AuthService) IsDuplicateEmail(email string) bool {
	res := s.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) (bool, error) {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false,  errors.New("wrong password")
	}
	return true, nil
}
