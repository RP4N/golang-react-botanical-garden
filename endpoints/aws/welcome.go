package main

import (
	"reimanexample/gardenService/api/welcome"
	"reimanexample/gardenService/internal/routing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

func main() {
	engine := routing.Build()

	// You may notice that this is identical to our localdev setup.
	// As you add more endpoints, you may wish to store the endpoints in a separate list, used by both localdev and aws.
	routing.AddRoute(engine, welcome.Path, welcome.Method, welcome.HandleRequest)

	proxy := func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		adapter := ginadapter.New(engine)
		return adapter.Proxy(req)
	}

	lambda.Start(proxy)
}
