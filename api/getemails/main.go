package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/mjourard/newsletter/api/pkg"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type DynamoEmailsTableItem struct {
	Email   string `json:"email"`
	Addedts int64  `json:"addedts"`
}
type RecipientEmail struct {
	Email        string `json:"email"`
	TimestampUTC string `json:"addedtsutc"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pkg.SetContext(ctx)
	logger := pkg.GetLogger()
	var buf bytes.Buffer

	sess, err := session.NewSession(&aws.Config{})
	svc := dynamodb.New(sess)

	emails := make([]*RecipientEmail, 0)
	proj := expression.NamesList(expression.Name("email"), expression.Name("addedts"))
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		return pkg.ErrorResponse(
			400,
			err,
			"Failed to build a new expression builder while trying to read from the table",
			"Error while trying construct the scanner request to scan the table: "+err.Error(),
		), nil
	}
	result, err := svc.Scan(&dynamodb.ScanInput{
		ConsistentRead:           aws.Bool(false),
		Limit:                    aws.Int64(500),
		ExpressionAttributeNames: expr.Names(),
		ProjectionExpression:     expr.Projection(),
		TableName:                aws.String(os.Getenv(pkg.EnvTable)),
	})

	if err != nil {
		return pkg.ErrorResponse(400, err, "Error while scanning the table", "Error: while trying to scan table: "+err.Error()), nil
	}

	if len(result.Items) == 0 {
		logger.Info("No items within table, returning 404")
		return pkg.ErrorResponse(404, nil, "", "No results found in table"), nil
	}

	for idx, i := range result.Items {
		item := DynamoEmailsTableItem{}
		err := dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			logger.WithError(err).Errorf("error while unmarshalling on idx %d", idx)
			continue
		}
		tm := time.Unix(item.Addedts, 0)
		email := &RecipientEmail{
			Email:        item.Email,
			TimestampUTC: tm.Format("2006-01-02 15:04:05 MST"),
		}
		emails = append(emails, email)
	}

	logger.Info("Finished scanning table, returning response")
	body, err := json.Marshal(emails)
	if err != nil {
		return pkg.ErrorResponse(400, err, "Error while unmarshalling after a successful table scan", "Unable to return table scan due to an error during data transformation"), nil
	}
	json.HTMLEscape(&buf, body)

	return pkg.Response(200, buf.String()), nil
}

func main() {
	lambda.Start(Handler)
}
