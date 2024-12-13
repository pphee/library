package use_case

import (
	"context"
	"github.com/pphee/library/model"
)

type UseCase struct {
	LibraryRepository LibraryRepository
}

type LibraryRepository interface {
	CreateBook(ctx context.Context, book model.Book) error
	GetBooks(ctx context.Context) ([]model.Book, error)
	GetBookByID(ctx context.Context, id uint) (model.Book, error)
	UpdateBook(ctx context.Context, book model.Book) error
	DeleteBook(ctx context.Context, id uint) error
	SearchBooks(ctx context.Context, title string, author string, category string) ([]model.Book, error)
	FetchBook(ctx context.Context, bookID uint) (*model.Book, error)
	GetMostBorrowedBooks(ctx context.Context, limit int) ([]model.Book, error)
}

func New(repo LibraryRepository) *UseCase {
	return &UseCase{
		LibraryRepository: repo,
	}
}
