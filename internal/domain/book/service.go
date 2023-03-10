package book

import (
	"ca-library-go/internal/adapters/api/book"
	"context"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) book.Service {
	return &service{storage: storage}
}

func (s *service) CreateBook(ctx context.Context, dto *CreateBookDTO) *Book {
	return nil
}

func (s *service) GetBookByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAllBook(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
