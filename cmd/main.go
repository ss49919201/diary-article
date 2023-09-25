package main

import (
	"context"

	"github.com/ss49919201/diary/internal/server"
	"github.com/ss49919201/diary/internal/sloghelper"
)

func main() {
	if err := server.Run(); err != nil {
		sloghelper.Fatal(context.Background(), "failed to run server")
	}
}
