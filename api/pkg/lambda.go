package pkg

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

const (
	EnvDeployment string = "ENV_DEPLOYMENT"
	EnvTable      string = "DYNAMO_TABLE"
)

func GetLogger(ctx context.Context) *log.Entry {
	lc, _ := lambdacontext.FromContext(ctx)
	return log.WithFields(log.Fields{
		"env":   os.Getenv(EnvDeployment),
		"table": os.Getenv(EnvTable),
		"request_id": lc.AwsRequestID,
	})
}
