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
	Email          string `json:"email"`
	AddedTimeStamp int64  `json:"addedts"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	pkg.SetContext(ctx)
	logger := pkg.GetLogger()
	logger.Debugf("Received body: %s", request.Body)
	sess, err := session.NewSession(&aws.Config{})
	if errResp := pkg.VerifyRequestParameters(request); errResp != nil {
		return *errResp, nil
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	logger.Info("Decoding addemail request")
	var email AddEmail
	err = json.Unmarshal([]byte(request.Body), &email)
	if err != nil {
		return pkg.ErrorResponse(400, err, "Unable to unmarshal request body into AddEmail struct", "Unable to decode JSON request"), nil
	}

	recipientAddEmail := &RecipientAddEmail{
		Email:          email.Email,
		AddedTimeStamp: time.Now().Unix(),
	}
	av, err := dynamodbattribute.MarshalMap(recipientAddEmail)

	if err != nil {
		return pkg.ErrorResponse(400, err, "Unable to marshal the RecipientAddEmail struct into something usable for dynamodb", "Unable to convert addemail to a savable item"), nil
	}

	condition := expression.AttributeNotExists(expression.Name("email"))
	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return pkg.ErrorResponse(400, err, "Unable to build conditional expression when adding item to email", "Error while trying construct the putitem request to insert into the table"), nil
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:                     av,
		TableName:                aws.String(os.Getenv(pkg.EnvTableRecipients)),
		ConditionExpression:      expr.Condition(),
		ExpressionAttributeNames: expr.Names(),
	}

	logger.Info("Adding email to table")
	_, err = svc.PutItem(input)

	if err != nil {
		bodyMsg := "Error: unable to add email to database. Reason unknown."
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				bodyMsg = fmt.Sprintf("Email '%s' is already registered in the database.", recipientAddEmail.Email)
			}
		}
		return pkg.ErrorResponse(400, err, fmt.Sprintf("Unable to add item to table. Condition: %s", *expr.Condition()), bodyMsg), nil
	}

	logger.Info("Successfully added email to table")
	return pkg.SuccessNoContentResponse(), nil
}

func main() {
	lambda.Start(Handler)
}
