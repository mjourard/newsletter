package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	EnvDeployment string = "ENV_DEPLOYMENT"
	EnvTable      string = "DYNAMO_TABLE"
)

type DynamoEmailsTableItem struct {
	Email string `json:"email"`
	Addedts int64 `json:"addedts"`
}
type RecipientEmail struct {
	Email          string `json:"email"`
	TimestampUTC string  `json:"addedtsutc"`
}

var ctx = log.WithFields(log.Fields{
	"func": "getemails",
	"env":  os.Getenv(EnvDeployment),
})

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer

	sess, err := session.NewSession(&aws.Config{})
	svc := dynamodb.New(sess)

	emails := make([]*RecipientEmail, 0)
	proj := expression.NamesList(expression.Name("email"), expression.Name("addedts"))
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		ctx.WithError(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error while trying construct the scanner request to scan the table: " + err.Error(),
		}, nil
	}
	result, err := svc.Scan(&dynamodb.ScanInput{
		ConsistentRead:           aws.Bool(false),
		Limit:                    aws.Int64(500),
		ExpressionAttributeNames: expr.Names(),
		ProjectionExpression:     expr.Projection(),
		TableName:                aws.String(os.Getenv(EnvTable)),
	})

	if err != nil {
		ctx.WithError(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body: "Error: while trying to scan table: " + err.Error(),
		}, nil
	}

	if len(result.Items) == 0 {
		ctx.Info("No items within table, returning 404")
		return events.APIGatewayProxyResponse{
			StatusCode:        404,
			Body:              "No results found in table",
			IsBase64Encoded:   false,
		}, nil
	}

	for idx, i := range result.Items {
		item := DynamoEmailsTableItem{}
		err := dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			ctx.WithFields(log.Fields{
				"msg": "error while unmarshalling",
				"idx": idx,
			}).WithError(err)
			continue
		}
		tm := time.Unix(item.Addedts, 0)
		email := &RecipientEmail{
			Email:        item.Email,
			TimestampUTC: tm.Format("2006-01-02 15:04:05 MST"),
		}
		emails = append(emails, email)
	}

	ctx.Info("Finished scanning table, returning response")
	body, err := json.Marshal(emails)
	if err != nil {
		ctx.WithError(err)
		return events.APIGatewayProxyResponse{
			StatusCode:        400,
			Body:              "Unable to return table scan due to an error during data transformation",
			IsBase64Encoded:   false,
		}, err
	}
	json.HTMLEscape(&buf, body)

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
