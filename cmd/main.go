package main

import (
	author3 "ca-library-go/internal/adapters/api/author"
	author2 "ca-library-go/internal/adapters/db/author"
	"ca-library-go/internal/composites"
	"ca-library-go/internal/domain/author"
)

func main() {

	authorComposite := composites.NewAuthorComposite()

	authorStorage := author2.NewStorage()
	authorService := author.NewService(authorStorage)
	authorHandler := author3.NewHandler(authorService)

}
