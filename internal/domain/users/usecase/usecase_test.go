package usecase_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/zuluapp/go-libs/pkg/libs/error"
	parameterstore "github.com/zuluapp/go-libs/pkg/libs/parameterstore/mock"
	rest "github.com/zuluapp/go-libs/pkg/libs/rest/mock"
	"github.com/zuluapp/go-libs/pkg/utils"

	"github.com/carlitos-zulu/lambda-http/internal/domain/users/entities"
	"github.com/carlitos-zulu/lambda-http/internal/domain/users/usecase"
	"github.com/carlitos-zulu/lambda-http/internal/infraestructure/dependencies"
)

func Test_NewUseCase(t *testing.T) {
	ass := assert.New(t)

	paramStore := parameterstore.NewMockService()
	paramStore.PatchGetParameter(utils.Pointer("test"), nil)

	restService := rest.NewRestMock()
	restService.PatchMakeGetRequest(http.StatusOK, []byte("{}"), nil)

	container := *dependencies.MockDependencies(restService)

	ass.NotNil(usecase.NewUseCase(container))
}

func Test_GetUserData(t *testing.T) {
	ass := assert.New(t)

	paramStore := parameterstore.NewMockService()
	paramStore.PatchGetParameter(utils.Pointer("test"), nil)

	restService := rest.NewRestMock()

	container := *dependencies.MockDependencies(restService)

	uc := usecase.NewUseCase(container)

	testError := fmt.Errorf("test error")

	tests := []struct {
		status int
		body   string
		error  error
		assert func(user *entities.User, err Wrapper)
	}{
		{
			status: http.StatusOK,
			body:   "{}",
			assert: func(user *entities.User, err Wrapper) {
				ass.Nil(err)
				ass.NotNil(user)
			},
		},
		{
			status: http.StatusInternalServerError,
			body:   "",
			error:  testError,
			assert: func(user *entities.User, err Wrapper) {
				ass.Nil(user)
				ass.Equal(testError.Error(), err.Error())
			},
		},
		{
			status: http.StatusBadRequest,
			body:   "",
			assert: func(user *entities.User, err Wrapper) {
				ass.Nil(user)
				ass.Equal(http.StatusBadRequest, err.Code().Status)
			},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", idx), func(t *testing.T) {
			restService.PatchMakeGetRequest(tt.status, []byte(tt.body), tt.error)

			tt.assert(uc.GetUserData(entities.RequestData{UserID: "123"}))
		})
	}
}
