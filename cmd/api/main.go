package main

import (
	"github.com/zuluapp/go-libs/pkg/libs/lambda/application"

	"github.com/carlitos-zulu/lambda-http/internal/domain/users/entities"
	"github.com/carlitos-zulu/lambda-http/internal/domain/users/handler"
	"github.com/carlitos-zulu/lambda-http/internal/infraestructure/dependencies"
)

func main() {
	application.StartHttp[handler.Handler, dependencies.Dependencies, entities.RequestData, *entities.User, dependencies.Container]()
}
