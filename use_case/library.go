package use_case

import (
	"context"
	"fmt"
	"github.com/pphee/library/model"
)

func (u *UseCase) CreateBook(ctx context.Context, book model.Book) error {
	return u.LibraryRepository.CreateBook(ctx, book)
}

func (u *UseCase) GetBooks(ctx context.Context) ([]model.Book, error) {
	return u.LibraryRepository.GetBooks(ctx)
}

func (u *UseCase) GetBookByID(ctx context.Context, id uint) (model.Book, error) {
	return u.LibraryRepository.GetBookByID(ctx, id)
}

func (u *UseCase) UpdateBook(ctx context.Context, book model.Book) error {
	return u.LibraryRepository.UpdateBook(ctx, book)
}

func (u *UseCase) DeleteBook(ctx context.Context, id uint) error {
	return u.LibraryRepository.DeleteBook(ctx, id)
}

func (u *UseCase) SearchBooks(ctx context.Context, title string, author string, category string) ([]model.Book, error) {
	return u.LibraryRepository.SearchBooks(ctx, title, author, category)
}

func (u *UseCase) BorrowBook(ctx context.Context, bookID uint) error {
	book, err := u.LibraryRepository.FetchBook(ctx, bookID)
	if err != nil {
		return err
	}

	if book.IsBorrowed {
		return fmt.Errorf("book is already borrowed")
	}

	book.IsBorrowed = true
	book.BorrowCount++

	return u.LibraryRepository.UpdateBook(ctx, *book)
}

func (u *UseCase) ReturnBook(ctx context.Context, bookID uint) error {
	book, err := u.LibraryRepository.FetchBook(ctx, bookID)
	if err != nil {
		return err
	}

	if !book.IsBorrowed {
		return fmt.Errorf("book is not borrowed")
	}

	book.IsBorrowed = false
	return u.LibraryRepository.UpdateBook(ctx, *book)
}

func (u *UseCase) GetMostBorrowedBooks(ctx context.Context, limit int) ([]model.Book, error) {
	return u.LibraryRepository.GetMostBorrowedBooks(ctx, limit)
}
