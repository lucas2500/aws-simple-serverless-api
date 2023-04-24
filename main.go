package main

import (
	"aws-simple-serverless-api/routes"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(routes.Hello)
}
