package handler

import (
	"net/http"

	"github.com/zuluapp/go-libs/pkg/libs/lambda/context"
	"github.com/zuluapp/go-libs/pkg/utils"

	"github.com/carlitos-zulu/lambda-http/internal/domain/users/entities"
	"github.com/carlitos-zulu/lambda-http/internal/domain/users/usecase"
	"github.com/carlitos-zulu/lambda-http/internal/infraestructure/dependencies"
)

type Handler struct{}

func (Handler) ProcessRequest(ctx context.LambdaContext[entities.RequestData, dependencies.Container]) (*entities.User, error) {
	var request entities.RequestData

	switch ctx.RequestMethod() {
	case http.MethodPost:
		request = ctx.Body()
	case http.MethodGet:
		request = entities.RequestData{
			UserID: ctx.QueryParams()["userId"],
		}
	}

	usecase := usecase.NewUseCase(ctx.Dependencies())

	user, err := usecase.GetUserData(request)
	if err != nil {
		return nil, err
	}

	user.Method = ctx.RequestMethod()

	return user, nil
}

func (Handler) GetTags() *utils.Map {
	return &utils.Map{}
}
