package usecase

import (
	"fmt"
	"net/http"

	"github.com/zuluapp/go-libs/pkg/libs/error"
	"github.com/zuluapp/go-libs/pkg/libs/rest"

	"github.com/carlitos-zulu/lambda-http/internal/domain/users/entities"
	"github.com/carlitos-zulu/lambda-http/internal/infraestructure/dependencies"
)

type (
	UseCase interface {
		GetUserData(data entities.RequestData) (*entities.User, error.Wrapper)
	}

	useCase struct {
		rest rest.Rest
	}
)

func NewUseCase(container dependencies.Container) UseCase {
	return useCase{
		rest: container.Rest(),
	}
}

func (uc useCase) GetUserData(data entities.RequestData) (*entities.User, error.Wrapper) {
	status, body, err := uc.rest.MakeGetRequest(fmt.Sprintf("/users/user/%s", data.UserID), nil)
	if err != nil {
		return nil, error.Wrap(err)
	}

	if status != http.StatusOK {
		return nil, error.ReturnWrappedErrorFromStatus(status, fmt.Errorf("call users error"))
	}

	user := &entities.User{}
	_ = body.GetStruct(user)

	return user, nil
}
