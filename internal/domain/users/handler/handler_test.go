package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zuluapp/go-libs/pkg/libs/lambda/context"
	restMock "github.com/zuluapp/go-libs/pkg/libs/rest/mock"

	"github.com/carlitos-zulu/lambda-http/internal/domain/users/entities"
	"github.com/carlitos-zulu/lambda-http/internal/domain/users/handler"
	"github.com/carlitos-zulu/lambda-http/internal/infraestructure/dependencies"
)

func Test_Handler(t *testing.T) {
	ass := assert.New(t)

	handler := handler.Handler{}

	tags := handler.GetTags()

	ass.NotNil(tags)

	rest := restMock.NewRestMock()

	container := dependencies.MockDependencies(rest)

	testError := fmt.Errorf("test error")

	tests := []struct {
		method string
		status int
		error  error
		query  map[string]string
		assert func(user *entities.User, err error)
	}{
		{
			method: http.MethodPost,
			status: http.StatusOK,
			assert: func(user *entities.User, err error) {
				ass.Nil(err)
				ass.NotNil(user)
			},
		},
		{
			method: http.MethodGet,
			status: http.StatusOK,
			query:  map[string]string{"userId": "123"},
			assert: func(user *entities.User, err error) {
				ass.Nil(err)
				ass.NotNil(user)
			},
		},
		{
			method: http.MethodPost,
			status: http.StatusInternalServerError,
			error:  testError,
			assert: func(user *entities.User, err error) {
				ass.Nil(user)
				ass.Equal(testError.Error(), err.Error())
			},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test_%v", idx), func(t *testing.T) {
			rest.PatchMakeGetRequest(tt.status, []byte(""), tt.error)

			ctx := context.InitCustomTestContext(context.TestContext[entities.RequestData, dependencies.Container]{
				Body: &entities.RequestData{
					UserID: "123",
				},
				Dependencies:  *container,
				RequestMethod: tt.method,
				Query:         tt.query,
			})

			tt.assert(handler.ProcessRequest(ctx))
		})
	}
}
