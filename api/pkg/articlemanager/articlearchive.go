package articlemanager

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

type ArticleManager struct {
	session         *session.Session
	dynamo          *dynamodb.DynamoDB
	s3              *s3.S3
	dynamoTableName *string
	cloudfrontURL   *string
}

func New(sess *session.Session, dynamoTableName *string, cloudfrontURL *string) ArticleManager {
	dynamoSvc := dynamodb.New(sess)
	s3Svc := s3.New(sess)
	return ArticleManager{
		session:         sess,
		dynamo:          dynamoSvc,
		s3:              s3Svc,
		dynamoTableName: dynamoTableName,
		cloudfrontURL:   cloudfrontURL,
	}
}

func (a *ArticleManager) AddNewArchivedArticle(article *ArchivesArchivedArticle, tableName string) error {
	av, err := dynamodbattribute.MarshalMap(article)

	if err != nil {
		return errors.Wrap(err, "Unable to marshal the ArchivesArchivedArticle struct into something usable for dynamodb")
	}

	condition := expression.AttributeNotExists(expression.Name("id"))
	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return errors.Wrap(err, "Unable to build conditional expression when adding article to archive")
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:                     av,
		TableName:                aws.String(tableName),
		ConditionExpression:      expr.Condition(),
		ExpressionAttributeNames: expr.Names(),
	}

	_, err = a.dynamo.PutItem(input)
	return err
}

func (a *ArticleManager) ListArchivedArticles(paginator *Paginator) ([]*ArchivesArchivedArticle, error) {
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("title"),
		expression.Name("img"),
		expression.Name("abstract"),
		expression.Name("addedts"),
		expression.Name("author"),
	)

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		return nil, errors.Wrap(err, "Got error while building expression")
	}

	output, err := a.dynamo.Scan(&dynamodb.ScanInput{
		ConsistentRead:            aws.Bool(true),
		ExclusiveStartKey:         paginator.lastEvaluatedKey,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		Limit:                     aws.Int64(paginator.getMax64()),
		ProjectionExpression:      expr.Projection(),
		TableName:                 a.dynamoTableName,
	})

	if err != nil {
		return nil, errors.Wrap(err, "Unable to scan table")
	}

	paginator.updatePaginator(output)

	if len(output.Items) == 0 {
		return []*ArchivesArchivedArticle{}, nil
	}

	articles := make([]*ArchivesArchivedArticle, 0)
	for idx, item := range output.Items {
		massagedItem := ArchivesScannedArchivedArticle{}
		err = dynamodbattribute.UnmarshalMap(item, &massagedItem)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Unable to unmarshal item w/ index # %d, id = %s", idx, *item["id"].S))
		}
		imgKey := a.GetImgKey(&massagedItem.Img)
		articles = append(articles, &ArchivesArchivedArticle{
			Id:             massagedItem.Id,
			Title:          massagedItem.Title,
			Img:            imgKey,
			Abstract:       massagedItem.Abstract,
			Author:         massagedItem.Author,
			AddedTimeStamp: int64(massagedItem.AddedTimeStamp),
		})
	}

	return articles, err
}

func (a *ArticleManager) GetImgKey(key *string) string {
	return *a.cloudfrontURL + "/" + *key
}

func CreateLastEvaluatedKey(lastID string) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"id": &dynamodb.AttributeValue{
			S: &lastID,
		},
	}
}
