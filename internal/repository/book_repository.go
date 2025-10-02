package repository

import (
	"BookStore/internal/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *domain.Book) error
	FindByID(id uint) (*domain.Book, error)
	FindByYear(year uint) (*domain.Book, error)
	SearchByTitle(title string) ([]domain.Book, error)
	SearchByAuthor(author string) ([]domain.Book, error)
	Update(book *domain.Book) error
	Delete(book *domain.Book) error
	ListAll() ([]domain.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *domain.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) FindByID(id uint) (*domain.Book, error) {
	var book domain.Book
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) FindByYear(year uint) (*domain.Book, error) {
	var book domain.Book
	if err := r.db.Where("year = ?", year).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) SearchByTitle(title string) ([]domain.Book, error) {
	var books []domain.Book
	if err := r.db.Where("title LIKE ?", title).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) SearchByAuthor(author string) ([]domain.Book, error) {
	var books []domain.Book
	if err := r.db.Where("author LIKE ?", author).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) Update(book *domain.Book) error {
	if err := r.db.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) Delete(book *domain.Book) error {
	if err := r.db.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) ListAll() ([]domain.Book, error) {
	var books []domain.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}