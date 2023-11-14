package main

import (
	"context"

	"github.com/ss49919201/diary-article/internal/dicontainer"
	"github.com/ss49919201/diary-article/internal/handler"
	"github.com/ss49919201/diary-article/internal/queryservice"
	"github.com/ss49919201/diary-article/internal/server"
	"github.com/ss49919201/diary-article/internal/sloghelper"
	"github.com/ss49919201/diary-article/internal/usecase"
)

func main() {
	dicontainer.Provide(handler.NewHealthCheck)
	dicontainer.Provide(handler.NewListArticles)
	dicontainer.Provide(usecase.NewListArticles)
	dicontainer.Provide(queryservice.NewListArticles)

	if err := server.Run(); err != nil {
		sloghelper.Fatal(context.Background(), "failed to run server")
	}
}
