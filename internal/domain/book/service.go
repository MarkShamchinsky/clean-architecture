package book

import (
	"ca-library-go/internal/adapters/api/book"
	"ca-library-go/internal/domain/author"
	"context"
)

type service struct {
	storage       Storage
	authorService author.Service
}

func NewService(storage Storage) book.Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *CreateBookDTO) *Book {

	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
