package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/mjourard/newsletter/api/pkg"
	log "github.com/sirupsen/logrus"
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

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	logger := pkg.GetLogger(ctx)
	logger.Debugf("Received body: %s", request.Body)
	sess, err := session.NewSession(&aws.Config{})
	if appType, ok := request.Headers["content-type"]; !ok || appType != "application/json"{
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Body:              "Error: request header was not 'application/json",
		}, nil
	}


	// Create DynamoDB client
	svc := dynamodb.New(sess)

	logger.Info("Decoding addemail request")
	var email AddEmail
	err = json.Unmarshal([]byte(request.Body), &email)
	if err != nil {
		logger.WithError(err).Error("Unable to unmarshal request body into AddEmail struct")
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
		logger.WithError(err).Error("Unable to marshal the RecipientAddEmail struct into something usable for dynamodb")
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Body:              "Error: unable to convert addemail to a savable item",
		}, nil
	}

	condition := expression.AttributeNotExists(expression.Name("email"))
	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		logger.WithError(err).Error("Unable to build conditional expression when adding item to email")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error while trying construct the putitem request to insert into the table",
		}, nil
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(os.Getenv(pkg.EnvTable)),
		ConditionExpression: expr.Condition(),
		ExpressionAttributeNames: expr.Names(),
	}

	logger.Info("Adding email to table")
	_, err = svc.PutItem(input)

	if err != nil {
		logger.WithError(err).Errorf("Unable to add item to table. Condition: %s", *expr.Condition())
		bodyMsg := "Error: unable to add email to database. Reason unknown."
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				bodyMsg = fmt.Sprintf("Error: email '%s' is already registered in the database.", recipientAddEmail.Email)
			}
		}
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Body:               bodyMsg,
		}, nil
	}

	logger.Info("Successfully added email to table")
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 204}, nil
}


func main() {
	log.SetOutput(os.Stdout)
	//TODO: set log output level -
	log.SetLevel(log.DebugLevel)
	lambda.Start(Handler)
}
