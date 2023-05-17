package main

import (
	. "context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
	. "github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/zuluapp/go-libs/pkg/constants"
	"github.com/zuluapp/go-libs/pkg/libs/lambda/application"
	"github.com/zuluapp/go-libs/pkg/libs/lambda/context"
	"github.com/zuluapp/go-libs/pkg/libs/metrics"
	parameterstore "github.com/zuluapp/go-libs/pkg/libs/parameterstore/mock"
	rest "github.com/zuluapp/go-libs/pkg/libs/rest/mock"
	"github.com/zuluapp/go-libs/pkg/libs/uuid"
	"github.com/zuluapp/go-libs/pkg/utils"

	"github.com/carlitos-zulu/lambda-http/internal/domain/users/entities"
	"github.com/carlitos-zulu/lambda-http/internal/infraestructure/dependencies"
)

type (
	testHttpHandler struct{}

	testLambdaHttpFactory[RQ any] struct {
		// property to mock behaviors
		body RQ
		ctx  any

		// properties to access the results
		response events.APIGatewayProxyResponse
		error    error
	}
)

func Test_Main(t *testing.T) {
	ass := assert.New(t)

	os.Setenv(constants.Application, "lambda-http")

	metrics.SetUpLambdaMock[entities.RequestData, dependencies.Container](nil)

	paramStore := parameterstore.NewMockService()

	paramStore.PatchGetParameter(utils.Pointer("test"), nil)

	restService := rest.NewRestMock()
	restService.PatchMakeGetRequest(http.StatusOK, []byte("{}"), nil)

	ctx := context.InitCustomTestContext(context.TestContext[entities.RequestData, dependencies.Container]{
		ParamStore:   paramStore,
		Dependencies: *dependencies.MockDependencies(restService),
	})

	recorder := &testLambdaHttpFactory[entities.RequestData]{
		body: entities.RequestData{
			UserID: "123",
		},
		ctx: ctx,
	}

	application.MockLambdaFactory(recorder)

	ass.NotPanics(main)

	ass.Nil(recorder.error)
	ass.Equal(http.StatusOK, recorder.response.StatusCode)
}

func (tt *testLambdaHttpFactory[RQ]) Start() func(handler interface{}) {
	return func(handler interface{}) {
		fn := handler.(func(ctx Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error))

		requestID := uuid.NewString()

		lambdaContext := &LambdaContext{
			AwsRequestID:       requestID,
			InvokedFunctionArn: "test",
		}

		ctxa := NewContext(Background(), lambdaContext)

		body, _ := json.Marshal(tt.body)

		event := events.APIGatewayProxyRequest{
			Path:       "/",
			HTTPMethod: http.MethodPost,
			Headers: map[string]string{
				"authorization": "test",
			},
			QueryStringParameters: map[string]string{},
			PathParameters:        map[string]string{},
			Body:                  string(body),
			RequestContext: events.APIGatewayProxyRequestContext{
				HTTPMethod: http.MethodPost,
			},
		}

		response, err := fn(ctxa, event)

		tt.response = response
		tt.error = err
	}
}

func (tt *testLambdaHttpFactory[RQ]) GetContext() any {
	return tt.ctx
}
