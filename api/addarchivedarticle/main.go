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

type AddArchivedArticle struct {
	Title    string `json:"title"`
	Img      string `json:"img"`
	Abstract string `json:"abstract"`
	Author   string `json:"author"`
}

type ArchivesArchivedArticle struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Img            string `json:"img"`
	Abstract       string `json:"abstract"`
	Author         string `json:"author"`
	AddedTimeStamp int64  `json:"addedts"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pkg.SetContext(ctx)
	logger := pkg.GetLogger()
	logger.Debugf("Received body: %s", request.Body)
	sess, err := session.NewSession(&aws.Config{})
	if appType, ok := request.Headers["content-type"]; !ok || appType != "application/json" {
		return pkg.ErrorResponse(400, nil, "", "Request header 'content-type' was not 'application/json'"), nil
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	logger.Info("Decoding addarchivedarticle request")
	var archivedArticle AddArchivedArticle
	err = json.Unmarshal([]byte(request.Body), &archivedArticle)
	if err != nil {
		return pkg.ErrorResponse(400, err, "Unable to unmarshal request body into AddArchivedArticle struct", "Unable to decode JSON request"), nil
	}

	//load the bucket from the environment
	bucket := os.Getenv("S3_BUCKET_ARTICLE_ASSETS")

	//upload the image to S3
	key, id, err := UploadToS3(sess, archivedArticle.Img, ContentTypeAutoDetect, bucket, PreviewContentPrefix)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to upload preview image to S3: %v", err)
		return pkg.ErrorResponse(500, err, errMsg, errMsg), nil
	}

	archivesArchivedArticle := &ArchivesArchivedArticle{
		Id:             id,
		Title:          archivedArticle.Title,
		Img:            key,
		Abstract:       archivedArticle.Abstract,
		Author:         archivedArticle.Author,
		AddedTimeStamp: time.Now().Unix(),
	}
	av, err := dynamodbattribute.MarshalMap(archivesArchivedArticle)

	if err != nil {
		return pkg.ErrorResponse(400, err, "Unable to marshal the ArchivesArchivedArticle struct into something usable for dynamodb", "Unable to convert addarchivearticle to a savable item"), nil
	}

	condition := expression.AttributeNotExists(expression.Name("id"))
	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return pkg.ErrorResponse(400, err, "Unable to build conditional expression when adding article to archive", "Error while trying construct the putitem request to insert into the table"), nil
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:                     av,
		TableName:                aws.String(os.Getenv(pkg.EnvTableArticleArchive)),
		ConditionExpression:      expr.Condition(),
		ExpressionAttributeNames: expr.Names(),
	}

	logger.Info("Adding article to archive table")
	_, err = svc.PutItem(input)

	if err != nil {
		bodyMsg := "Error: unable to add article to archive database. Reason unknown."
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				bodyMsg = fmt.Sprintf("id '%s' is already registered in the database.", archivesArchivedArticle.Id)
			}
		}
		return pkg.ErrorResponse(400, err, fmt.Sprintf("Unable to add item to table. Condition: %s", *expr.Condition()), bodyMsg), nil
	}

	logger.Info("Successfully added article to archive")
	return pkg.SuccessObjectResponse(archivesArchivedArticle), nil
}

func main() {
	lambda.Start(Handler)
}
