package composites

import (
	"ca-library-go/internal/adapters/api"
	book2 "ca-library-go/internal/adapters/api/book"
	book3 "ca-library-go/internal/adapters/db/book"
	"ca-library-go/internal/domain/book"
)

type BookComposite struct {
	Storage book.Storage
	Service book2.Service
	Handler api.Handler
}

func NewBookComposite() (*BookComposite, error) {
	storage := book3.NewStorage()
	service := book.NewService(storage)
	handler := book2.NewHandler(service)

	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
