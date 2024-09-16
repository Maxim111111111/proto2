package main

import (
	"Pr_grpc_go/sso/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}