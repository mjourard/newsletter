package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	EnvDeployment          string = "ENV_DEPLOYMENT"
	EnvTableRecipients     string = "DYNAMO_TABLE_RECIPIENTS"
	EnvTableArticleArchive string = "DYNAMO_TABLE_ARTICLE_ARCHIVE"
	EnvLogLevel            string = "LOG_LEVEL"
)

var (
	fields log.Fields = log.Fields{
		"env": os.Getenv(EnvDeployment),
	}
)

func SetContext(ctx context.Context) {
	lc, _ := lambdacontext.FromContext(ctx)
	fields = log.Fields{
		"env":        os.Getenv(EnvDeployment),
		"table":      os.Getenv(EnvTableRecipients),
		"request_id": lc.AwsRequestID,
	}
	log.SetOutput(os.Stdout)
	SetLogLevel()
}

//SetLogLevel will parse the environment variable EnvLogLevel into the logger's level and
//set it as such. It is Case Insensitive. The possible logging values are:
//panic
//fatal
//error
//warn
//info
//debug
//trace
func SetLogLevel() {
	lvlStr := os.Getenv(EnvLogLevel)
	if lvlStr == "" {
		lvlStr = "info"
	}
	lvl, err := log.ParseLevel(lvlStr)
	if err != nil {
		lvl = log.InfoLevel
	}
	log.SetLevel(lvl)
}

func GetLogger() *log.Entry {
	return log.WithFields(fields)
}

//TODO: if additional properties of the APIGatewayProxyResponse needs to be set, add a new method that takes in one and appends/updates values

func ErrorResponse(status int, err error, loggableMsg string, errorMessage string) events.APIGatewayProxyResponse {
	if err != nil || len(loggableMsg) > 0 {
		GetLogger().WithError(err).Error(loggableMsg)
	}

	var buf bytes.Buffer
	body, err := json.Marshal(map[string]string{"message": errorMessage})
	if err != nil {
		body = handleMarshallingError(err)
	}
	json.HTMLEscape(&buf, body)
	return Response(status, buf.String())
}

func SuccessNoContentResponse() events.APIGatewayProxyResponse {
	return Response(204, "")
}

func SuccessObjectResponse(bodyObj interface{}) events.APIGatewayProxyResponse {
	body, err := json.Marshal(bodyObj)
	if err != nil {
		body = handleMarshallingError(err)
	}
	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)
	return Response(200, buf.String())
}

func Response(status int, body string) events.APIGatewayProxyResponse {
	//TODO: add a method for checking a whitelist of origins to send back (can only send back one) so we don't need to send back a '*'
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "POST, GET, OPTIONS",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}

func handleMarshallingError(err error) []byte {
	GetLogger().WithError(err).Error("Unable to json marshal the error message: %v", err)
	return []byte(`{"message": "an unexpected error occurred while processing"}`)
}
