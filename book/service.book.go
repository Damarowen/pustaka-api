package book

import (
	"errors"
	"fmt"
	"log"
	"pustaka-api/dto"
	"pustaka-api/models"

	"github.com/mashingan/smapping"
)

type Iservice interface {
	FindAll() ([]models.Book, error)
	FindById(ID uint) (models.Book, error)
	Create(book dto.BookRequest) (models.Book, error)
	Update(book dto.BookUpdateDTO) (models.Book, error)
	Delete(book models.Book) (models.Book, error)
	IsAllowedToEdit(userID string, bookID uint) (bool, error)
}

type Service struct {
	PustakaApiRepository Irepository
}

func NewBookService(bookRepo Irepository) Iservice {
	return &Service{PustakaApiRepository: bookRepo}
}

func (s *Service) FindAll() ([]models.Book, error) {
	books, err := s.PustakaApiRepository.FindAll()
	return books, err
}

func (s *Service) FindById(ID uint) (models.Book, error) {

	book, err := s.PustakaApiRepository.FindById(ID)

	return book, err

}

func (s *Service) Create(bookRequest dto.BookRequest) (models.Book, error) {

	book := models.Book{}

	//* BEFORE
	// book := models.Book{
	// 	Title:       bookRequest.Title,
	// 	Price:       bookRequest.Price,
	// 	Description: bookRequest.Description,
	// 	Rating:      bookRequest.Rating,
	// }

	//* mappong dengan lib
	err := smapping.FillStruct(&book, smapping.MapFields(&bookRequest))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	newBook, err := s.PustakaApiRepository.Create(book)
	return newBook, err
}

func (s *Service) Update(b dto.BookUpdateDTO) (models.Book, error) {

	// find, err := s.PustakaApiRepository.FindById(ID)

	// if find.ID == 0 {
	// 	return find, errors.New("ID NOT FOUND")
	// }

	book := models.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res, err := s.PustakaApiRepository.Update(book)
	return res, err
}

func (s *Service) Delete(book models.Book) (models.Book, error) {

	b, err := s.PustakaApiRepository.Delete(book)
	return b, err
}

func (service *Service) IsAllowedToEdit(userID string, bookID uint) (bool, error) {
	b, _ := service.PustakaApiRepository.FindById(bookID)
	if b.ID == 0 {
		return false, errors.New("ID NOT FOUND")
	}
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id, nil
}
