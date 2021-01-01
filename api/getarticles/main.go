package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mjourard/newsletter/api/pkg"
	"github.com/mjourard/newsletter/api/pkg/articlemanager"
	"os"
)

type ArticlePages struct {
	Max    int    `json:"max"`
	LastID string `json:"lastid"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pkg.SetContext(ctx)
	logger := pkg.GetLogger()
	var buf bytes.Buffer

	if errResp := pkg.VerifyRequestParameters(request); errResp != nil {
		return *errResp, nil
	}

	sess, err := session.NewSession(&aws.Config{})
	var pages ArticlePages
	if errResp := pkg.DecodeRequest(request, &pages); errResp != nil {
		return *errResp, nil
	}

	artManager := articlemanager.New(sess, aws.String(os.Getenv(pkg.EnvTableArticleArchive)), aws.String(os.Getenv(pkg.EnvCloudfrontURL)))
	var lastKey map[string]*dynamodb.AttributeValue
	if pages.LastID != "" {
		lastKey = articlemanager.CreateLastEvaluatedKey(pages.LastID)
	}
	paginator := articlemanager.NewPaginator(pages.Max, lastKey)

	logger.Info("Beginning table scan")
	articles, err := artManager.ListArchivedArticles(paginator)
	if err != nil {
		return pkg.ErrorResponse(400, err, err.Error(), "Unable to retrieve a list of the archived articles."), nil
	}

	logger.Info("Finished scanning table, returning response")
	body, err := json.Marshal(articles)
	if err != nil {
		return pkg.ErrorResponse(400, err, "Error while unmarshalling after a successful table scan", "Unable to return table scan due to an error during data transformation"), nil
	}
	json.HTMLEscape(&buf, body)

	return pkg.Response(200, buf.String()), nil
}

func main() {
	lambda.Start(Handler)
}
