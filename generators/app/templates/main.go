package main

import (
	"bitbucket.org/credomobile/<%= lambdaName %>/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.EventHandler)
}
