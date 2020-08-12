package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"time"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type AddEmail struct {
	Email string `json:"email"`
}

type RecipientAddEmail struct {
	Email string `json:"email"`
	AddedTimeStamp int64 `json:"addedts"`
}

const (
	EnvDeployment string = "ENV_DEPLOYMENT"
	EnvTable string = "DYNAMO_TABLE"
)

var ctx = log.WithFields(log.Fields{
	"func": "addemail",
	"env": os.Getenv(EnvDeployment),
})

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ctx.Debugf("Received body: %s", request.Body)
	sess, err := session.NewSession(&aws.Config{})
	for header, value := range request.Headers {
		ctx.Debugf("%s: %s", header, value)
	}

	if appType, ok := request.Headers["content-type"]; !ok || appType != "application/json"{
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Body:              "Error: request header was not 'application/json",
		}, nil
	}


	// Create DynamoDB client
	svc := dynamodb.New(sess)

	ctx.Info("Decoding addemail request")
	var email AddEmail
	err = json.Unmarshal([]byte(request.Body), &email)
	if err != nil {
		ctx.WithError(err)
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Body:              "Error: unable to decode JSON request",
		}, nil
	}

	recipientAddEmail := &RecipientAddEmail{
		Email:          email.Email,
		AddedTimeStamp: time.Now().Unix(),
	}
	av, err := dynamodbattribute.MarshalMap(recipientAddEmail)

	if err != nil {
		ctx.WithError(err)
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Body:              "Error: unable to convert addemail to a savable item",
		}, nil
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(os.Getenv(EnvTable)),
	}

	ctx.Info("Adding email to table")
	_, err = svc.PutItem(input)

	if err != nil {
		ctx.WithError(err)
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Body:              "Error: unable to add item to database",
		}, nil
	}

	ctx.Info("Successfully added email to table")
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 204}, nil
}


func main() {
	log.SetOutput(os.Stdout)
	//TODO: set log output level -
	log.SetLevel(log.DebugLevel)
	lambda.Start(Handler)
}
