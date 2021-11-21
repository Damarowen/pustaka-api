package book

import "gorm.io/gorm"

type Irepository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
}

type PustakaApiRepository struct {
	pustaka_api *gorm.DB
}

func NewRepository(db *gorm.DB) *PustakaApiRepository {
	return &PustakaApiRepository{db}
}

func (r *PustakaApiRepository ) FindAll() ([]Book, error){
	var books []Book

	err := r.pustaka_api.Find(&books).Error

	return books, err
}

func (r *PustakaApiRepository ) FindById(ID int) (Book, error){
	var book  Book

	err := r.pustaka_api.Find(&book, ID).Error

	return book, err
}

func (r *PustakaApiRepository ) Create(book Book) (Book, error){

	err := r.pustaka_api.Create(&book).Error

	return book, err
}