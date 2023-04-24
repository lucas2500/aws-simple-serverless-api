package routes

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func Hello(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var person Person

	err := json.Unmarshal([]byte(request.Body), &person)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	ResponseBody := map[string]string{
		"Message": "Request processed successfully",
		"Data":    request.Body,
	}

	jbytes, err := json.Marshal(&ResponseBody)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jbytes),
	}

	return response, nil
}
