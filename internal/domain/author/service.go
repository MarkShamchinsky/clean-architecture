package author

import (
	"ca-library-go/internal/adapters/api/author"
	"context"
)

type Service struct {
	storage Storage
}

func NewService(storage Storage) author.Service {
	return &Service{storage: storage}
}

func (s *Service) Create(ctx context.Context, dto *CreateAuthorDTO) *Author {
	return nil
}

func (s *Service) GetByUUID(ctx context.Context, uuid string) *Author {
	return s.storage.GetOne(uuid)
}

func (s *Service) GetAll(ctx context.Context, limit, offset int) []*Author {
	return s.storage.GetAll(limit, offset)
}
