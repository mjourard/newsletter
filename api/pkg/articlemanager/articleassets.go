package articlemanager

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"net/http"
	"strings"
)

type AssetsManager struct {
	session    *session.Session
	s3         *s3.S3
	bucketName *string
}

const (
	PreviewContentPrefix string = "preview/"
)

func NewAssetsManager(sess *session.Session, assetBucketName *string) AssetsManager {
	s3Svc := s3.New(sess)
	return AssetsManager{
		session:    sess,
		s3:         s3Svc,
		bucketName: assetBucketName,
	}
}

//UploadImageToS3 will upload the passed in content to the specified bucket with the optional bucketPrefix
//The key generated is a random uuid from the segmentio/ksuid library. The string returned is the key
//in which the content is saved under
func (a *AssetsManager) UploadImageToS3(sess *session.Session, content string, bucketPrefix string) (string, string, error) {
	svc := s3.New(sess)
	if len(content) == 0 {
		return "", "", errors.New("empty content passed in while trying to upload to S3")
	}
	contentTypeHeader := GetContentTypeStr(content)
	//generate the key to upload
	if len(bucketPrefix) > 0 {
		//ensure it ends in a single forward slash
		if bucketPrefix[len(bucketPrefix)-1:] != "/" {
			bucketPrefix += "/"
		}
		if bucketPrefix == "/" {
			bucketPrefix = ""
		}
	}
	id := ksuid.New().String()
	key := fmt.Sprintf("%s%s", bucketPrefix, id)

	parsedContent, err := a.parseImgContent(content)
	if err != nil {
		return "", "", fmt.Errorf("unable to parse the image content before uploading to s3: %w", err)
	}

	//upload to s3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:          bytes.NewReader(parsedContent),
		ACL:           aws.String("bucket-owner-full-control"),
		Bucket:        a.bucketName,
		ContentLength: aws.Int64(int64(len(parsedContent))),
		ContentType:   &contentTypeHeader,
		Key:           &key,
		Metadata:      nil, //possibly need to be setting this
	})
	if err != nil {
		return "", "", fmt.Errorf("upload to s3 failed: %w", err)
	}

	return key, id, nil
}

func GetContentTypeStr(content string) string {
	buff := []byte(content[:512])
	return http.DetectContentType(buff)
}

func (a *AssetsManager) parseImgContent(content string) ([]byte, error) {
	s := strings.SplitN(content, ",", 2)
	if len(s) != 2 {
		return nil, errors.New("unable to parse image content: no comma found to split on")
	}
	decoded, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse image content: failed to base64 decode the image content")
	}
	return decoded, nil
}
