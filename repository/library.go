package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/pphee/library/model"
	"github.com/pphee/library/use_case"
	"gorm.io/gorm"
	"strings"
)

type libraryRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) use_case.LibraryRepository {
	return &libraryRepo{
		DB: db,
	}
}

func (r *libraryRepo) CreateBook(ctx context.Context, book model.Book) error {
	return r.DB.WithContext(ctx).Create(&book).Error
}

func (r *libraryRepo) GetBooks(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	err := r.DB.WithContext(ctx).Find(&books).Error
	return books, err
}

func (r *libraryRepo) GetBookByID(ctx context.Context, id uint) (model.Book, error) {
	var book model.Book
	err := r.DB.WithContext(ctx).First(&book, id).Error
	return book, err
}

func (r *libraryRepo) UpdateBook(ctx context.Context, book model.Book) error {
	return r.DB.WithContext(ctx).Save(&book).Error
}

func (r *libraryRepo) DeleteBook(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&model.Book{}, id).Error
}

func (r *libraryRepo) SearchBooks(ctx context.Context, title string, author string, category string) ([]model.Book, error) {
	var books []model.Book
	query := r.DB.WithContext(ctx).Model(&model.Book{})

	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}
	if author != "" {
		query = query.Where("LOWER(author) LIKE ?", "%"+strings.ToLower(author)+"%")
	}
	if category != "" {
		query = query.Where("LOWER(category) LIKE ?", "%"+strings.ToLower(category)+"%")
	}

	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *libraryRepo) FetchBook(ctx context.Context, bookID uint) (*model.Book, error) {
	var book model.Book
	if err := r.DB.WithContext(ctx).First(&book, bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}
	return &book, nil
}

func (r *libraryRepo) GetMostBorrowedBooks(ctx context.Context, limit int) ([]model.Book, error) {
	var books []model.Book

	if err := r.DB.WithContext(ctx).
		Model(&model.Book{}).
		Order("borrow_count DESC").
		Limit(limit).
		Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
