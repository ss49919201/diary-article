package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ss49919201/diary-article/internal/dicontainer"
	"github.com/ss49919201/diary-article/internal/env"
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

	if err := server.NewServer(env.Address()).ListenAndServe(); err != nil {
		sloghelper.Fatal(context.Background(), fmt.Sprintf("failed to run server: %s", err))
		os.Exit(1)
	}
}
