package storage

import (
	"github.com/Flamiry/g2-books/internal/domain/models"
	"fmt"
)

type MemStorage struct {
	usersMap map[string]models.User
	booksMap map[string]models.Book
}

