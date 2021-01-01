package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mjourard/newsletter/api/pkg"
	"github.com/mjourard/newsletter/api/pkg/articlemanager"
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

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pkg.SetContext(ctx)
	logger := pkg.GetLogger()
	logger.Debugf("Received body: %s", request.Body)
	sess, err := session.NewSession(&aws.Config{})
	if errResp := pkg.VerifyRequestParameters(request); errResp != nil {
		return *errResp, nil
	}

	articleManager := articlemanager.New(sess, aws.String(os.Getenv(pkg.EnvTableArticleArchive)), aws.String(os.Getenv(pkg.EnvCloudfrontURL)))

	logger.Info("Decoding addarchivedarticle request")
	var archivedArticle AddArchivedArticle
	errResp := pkg.UnmarshalRequest(request, &archivedArticle)
	if errResp != nil {
		return *errResp, nil
	}

	//load the bucket from the environment
	bucket := os.Getenv(pkg.EnvS3BucketArticleAssets)

	//upload the image to S3
	key, id, err := UploadToS3(sess, archivedArticle.Img, ContentTypeAutoDetect, bucket, PreviewContentPrefix)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to upload preview image to S3: %v", err)
		return pkg.ErrorResponse(500, err, errMsg, errMsg), nil
	}

	archivesArchivedArticle := &articlemanager.ArchivesArchivedArticle{
		Id:             id,
		Title:          archivedArticle.Title,
		Img:            key,
		Abstract:       archivedArticle.Abstract,
		Author:         archivedArticle.Author,
		AddedTimeStamp: time.Now().Unix(),
	}
	err = articleManager.AddNewArchivedArticle(archivesArchivedArticle, os.Getenv(pkg.EnvTableArticleArchive))

	if err != nil {
		bodyMsg := err.Error()
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				bodyMsg = fmt.Sprintf("id '%s' is already registered in the database. Internal error: %s", archivesArchivedArticle.Id, err.Error())
			}
		}
		return pkg.ErrorResponse(400, err, bodyMsg, "Unable to convert addarchivearticle to a savable item"), nil
	}

	logger.Info("Successfully added article to archive")
	return pkg.SuccessObjectResponse(archivesArchivedArticle), nil
}

func main() {
	lambda.Start(Handler)
}
