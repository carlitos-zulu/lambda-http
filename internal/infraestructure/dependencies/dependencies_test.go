package dependencies_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zuluapp/go-libs/pkg/libs/lambda/context"
	parameterstore "github.com/zuluapp/go-libs/pkg/libs/parameterstore/mock"
	"github.com/zuluapp/go-libs/pkg/utils"

	"github.com/carlitos-zulu/lambda-http/internal/domain/users/entities"
	"github.com/carlitos-zulu/lambda-http/internal/infraestructure/dependencies"
)

func Test_Dependencies(t *testing.T) {
	ass := assert.New(t)

	deps := dependencies.Dependencies{}

	paramStore := parameterstore.NewMockService()

	paramStore.PatchGetParameter(utils.Pointer("test"), nil)

	ctx := context.InitCustomTestContext(context.TestContext[entities.RequestData, dependencies.Container]{
		ParamStore: paramStore,
	})

	container := deps.Start(ctx)

	ass.NotNil(container)
	ass.NotNil(container.Rest())
}

func Test_MockDependencies(t *testing.T) {
	ass := assert.New(t)

	ass.NotNil(dependencies.MockDependencies(nil))
}
