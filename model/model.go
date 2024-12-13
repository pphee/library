package model

type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Author      string `gorm:"not null"`
	Category    string
	BorrowCount int  `gorm:"default:0"`
	IsBorrowed  bool `gorm:"default:false"`
}

type BorrowRequest struct {
	BookID uint `json:"book_id"`
}
