package main

import (
	"context"

	"github.com/ss49919201/diary/internal/dicontainer"
	"github.com/ss49919201/diary/internal/handler"
	"github.com/ss49919201/diary/internal/queryservice"
	"github.com/ss49919201/diary/internal/server"
	"github.com/ss49919201/diary/internal/sloghelper"
	"github.com/ss49919201/diary/internal/usecase"
)

func main() {
	dicontainer.Provide(handler.NewHealthCheck)
	dicontainer.Provide(handler.NewListArticle)
	dicontainer.Provide(usecase.NewListArticles)
	dicontainer.Provide(queryservice.NewListArticles)

	if err := server.Run(); err != nil {
		sloghelper.Fatal(context.Background(), "failed to run server")
	}
}
