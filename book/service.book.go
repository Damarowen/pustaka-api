package book


type Iservice interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
}

type Service struct {
	PustakaApiRepository Irepository
}

func NewService(PustakaApiRepository Irepository) *Service {
	return &Service{PustakaApiRepository}
}

func (s *Service) FindAll() ([]Book, error) {
	books, err := s.PustakaApiRepository.FindAll()
	return books, err
}

func (s *Service) FindById(ID int) (Book, error) {

	book, err := s.PustakaApiRepository.FindById(ID)

	return book, err
}

func (s *Service) Create(bookRequest BookRequest) (Book, error) {

	book := Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}

	newBook, err := s.PustakaApiRepository.Create(book)
	return newBook, err
}
