package dependencies

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/carlitos-zulu/lambda-http/internal/domain/users/entities"

	"github.com/zuluapp/go-libs/pkg/libs/lambda/context"
	"github.com/zuluapp/go-libs/pkg/libs/rest"
)

type (
	Dependencies struct{}

	Container struct {
		rest rest.Rest
	}

	Params struct {
		PrivateApiUrlBase   string
		TraceabilityHeaders http.Header
	}
)

func (deps Dependencies) Start(ctx context.LambdaContext[entities.RequestData, Container]) Container {
	return Container{
		rest: getRestService(Params{
			PrivateApiUrlBase:   ctx.PrivateAPIUrlBase(),
			TraceabilityHeaders: ctx.TraceabilityHeaders(),
		}),
	}
}

func (container Container) Rest() rest.Rest {
	return container.rest
}

func getRestService(params Params) rest.Rest {
	service := rest.NewRestService(rest.ServiceConfig{
		BaseURL:             params.PrivateApiUrlBase,
		MaxIdleConnsPerHost: 50,
		RequestConfig: &rest.RequestConfig{
			DisableTimeout: false,
			Timeout:        3 * time.Second,
			ConnectTimeout: 1500 * time.Millisecond,
		},
	})

	service.SetDefaultHeaders(params.TraceabilityHeaders)
	service.SetTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	fmt.Println(`Rest configured with base:`, params.PrivateApiUrlBase)
	fmt.Println("Traceability headers:", params.TraceabilityHeaders)

	return service
}
