package service

import (
	"BookStore/internal/domain"
	"BookStore/internal/repository"
	"errors"
)

type BookService interface {
	Create(title, author string, year uint) (*domain.Book, error)
	GetByID(id uint) (*domain.Book, error)
	Search(title, author string, year uint) ([]domain.Book, error)
	Update(book *domain.Book) error
	Delete(book *domain.Book) error
	Paginate(page, pageSize int) ([]domain.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookService {
	return &bookService{repo: r}
}

func (s *bookService) Create(title, author string, year uint) (*domain.Book, error) {
	book := &domain.Book{
		Title:  title,
		Author: author,
		Year:   year,
	}
	check, _ := s.repo.FindByTitle(title)
	if check != nil {
		return nil, errors.New("book with this title already exists")
	}
	if  len(author) <  4 {
		return nil, errors.New("invalid book details")
	}
	return book, s.repo.Create(book)
}

func (s *bookService) GetByID(id uint) (*domain.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) Search(title, author string, year uint) ([]domain.Book, error) {
	bookMap := make(map[uint]domain.Book) 
	var result []domain.Book

	if title != "" {
		books, err := s.repo.SearchByTitle(title)
		if err != nil {
			return nil, err
		}
		for _, b := range books {
			bookMap[b.ID] = b
		}
	}

	if author != "" {
		books, err := s.repo.SearchByAuthor(author)
		if err != nil {
			return nil, err
		}
		for _, b := range books {
			bookMap[b.ID] = b
		}
	}

	if year != 0 {
		books, err := s.repo.SearchByYear(year)
		if err != nil {
			return nil, err
		}
		for _, b := range books {
			bookMap[b.ID] = b
		}
	}

	for _, b := range bookMap {
		result = append(result, b)
	}

	return result, nil
}



func (s *bookService) Update(book *domain.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) Delete(book *domain.Book) error {
	return s.repo.Delete(book)
}

func (s *bookService) Paginate(page, pageSize int) ([]domain.Book, error) {
	return s.repo.Paginate(page, pageSize)
}
