package book

import (
	"errors"
	"pustaka-api/models"
	"pustaka-api/dto"
)

type Iservice interface {
	FindAll() ([]models.Book, error)
	FindById(ID uint) (models.Book, error)
	Create(book dto.BookRequest) (models.Book, error)
	Update(ID  dto.BookRequest) (models.Book, error)
	Delete(ID  dto.BookRequest) (models.Book, error)
}

type Service struct {
	PustakaApiRepository Irepository
}

func NewService(PustakaApiRepository Irepository) *Service {
	return &Service{PustakaApiRepository}
}

func (s *Service) FindAll() ([]models.Book, error) {
	books, err := s.PustakaApiRepository.FindAll()
	return books, err
}

func (s *Service) FindById(ID uint) (models.Book, error) {

	book, err := s.PustakaApiRepository.FindById(ID)

	if book.ID == 0 {
		return book, errors.New("ID NOT FOUND")
	}

	return book, err

}

func (s *Service) Create(bookRequest  dto.BookRequest) (models.Book, error) {

	book := models.Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}

	newBook, err := s.PustakaApiRepository.Create(book)
	return newBook, err
}

func (s *Service) Update(ID uint, bookRequest  dto.BookRequest) (models.Book, error) {

	find, err := s.PustakaApiRepository.FindById(ID)

	if find.ID== 0 {
		return find, errors.New("ID NOT FOUND")
	}

	find.Title = bookRequest.Title
	find.Price = bookRequest.Price
	find.Description = bookRequest.Description
	find.Rating = bookRequest.Rating

	b, err := s.PustakaApiRepository.Update(find)
	return b, err
}

func (s *Service) Delete(ID uint) (models.Book, error) {

	find, err := s.PustakaApiRepository.FindById(ID)

	if find.ID== 0 {
		return find, errors.New("ID NOT FOUND")
	}

	b, err := s.PustakaApiRepository.Delete(find)
	return b, err
}
