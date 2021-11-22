package user

import (
	"log"
	"pustaka-api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InsertUser(user models.User) (models.User)
	UpdateUser(user models.User) models.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(userID string) models.User
}

type PustakaApiRepository struct {
	pustaka_api *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &PustakaApiRepository{
		pustaka_api: db,
	}
}

func (r *PustakaApiRepository) InsertUser(user models.User) models.User {
	user.Password = hashAndSalt([]byte(user.Password))
	r.pustaka_api.Save(&user)

	return user
}


func (r *PustakaApiRepository) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		r.pustaka_api.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	r.pustaka_api.Save(&user)
	return user
}

func (r *PustakaApiRepository) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := r.pustaka_api.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (r *PustakaApiRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return r.pustaka_api.Where("email = ?", email).Take(&user)
}

func (r *PustakaApiRepository) FindByEmail(email string) models.User {
	var user models.User
	r.pustaka_api.Where("email = ?", email).Take(&user)
	return user
}

func (r *PustakaApiRepository) ProfileUser(userID string) models.User {
	var user models.User
	r.pustaka_api.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}


func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
