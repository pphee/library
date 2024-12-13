package handler

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST("/books", h.CreateBook)
	router.GET("/books", h.GetBooks)
	router.GET("/books/:id", h.GetBookByID)
	router.PUT("/books/:id", h.UpdateBook)
	router.DELETE("/books/:id", h.DeleteBook)
	router.GET("/search", h.SearchBooks)
	router.POST("/books/:id/return", h.ReturnBook)
	router.POST("/books/:id/borrow", h.BorrowBook)
	router.GET("/most_borrowed", h.GetMostBorrowedBooks)
}
