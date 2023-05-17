package dependencies

import "github.com/zuluapp/go-libs/pkg/libs/rest"

func MockDependencies(rest rest.Rest) *Container {
	return &Container{rest}
}
