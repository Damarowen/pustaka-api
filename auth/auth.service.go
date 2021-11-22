package auth

import (
	"log"
	"pustaka-api/dto"
	"pustaka-api/models"
	"pustaka-api/user"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type IAuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	FindByEmail(email string) models.User
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

func (s *AuthService) VerifyCredential(email string, password string) interface{} {
	res := s.userRepository.VerifyCredential(email, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
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

func (s *AuthService) FindByEmail(email string) models.User {
	return s.userRepository.FindByEmail(email)
}

func (s *AuthService) IsDuplicateEmail(email string) bool {
	res := s.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
