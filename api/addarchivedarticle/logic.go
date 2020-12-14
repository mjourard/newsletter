package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/segmentio/ksuid"
	"net/http"
)

type ContentType int

const (
	ContentTypeBitmap ContentType = iota
	ContentTypeGif
	ContentTypeIcon
	ContentTypeJpeg
	ContentTypePng
	ContentTypeSvg
	ContentTypeTiff
	ContentTypeWebp
	ContentTypeAutoDetect
)

const (
	PreviewContentPrefix string = "preview/"
)

//UploadToS3 will upload the passed in content to the specified bucket with the optional bucketPrefix
//The key generated is a random uuid from the segmentio/ksuid library. The string returned is the key
//in which the content is saved under
func UploadToS3(sess *session.Session, content string, contentType ContentType, bucket string, bucketPrefix string) (string, string, error) {
	svc := s3.New(sess)
	if len(content) == 0 {
		return "", "", errors.New("empty content passed in while trying to upload to S3")
	}
	contentTypeHeader := GetContentTypeStr(contentType, content)
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

	//upload to s3
	_, err := svc.PutObject(&s3.PutObjectInput{
		Body:          bytes.NewReader([]byte(content)),
		ACL:           aws.String("bucket-owner-full-control"),
		Bucket:        &bucket,
		ContentLength: aws.Int64(int64(len(content))),
		ContentType:   &contentTypeHeader,
		Key:           &key,
		Metadata:      nil, //possibly need to be setting this
	})
	if err != nil {
		return "", "", fmt.Errorf("upload to s3 failed: %w", err)
	}

	return key, id, nil
}

func GetContentTypeStr(contentType ContentType, content string) string {
	switch contentType {
	case ContentTypeBitmap:
		return "image/bmp"
	case ContentTypeGif:
		return "image/gif"
	case ContentTypeIcon:
		return "image/vnd.microsoft.icon"
	case ContentTypeJpeg:
		return "image/jpeg"
	case ContentTypePng:
		return "image/png"
	case ContentTypeSvg:
		return "image/svg+xml"
	case ContentTypeTiff:
		return "image/tiff"
	case ContentTypeWebp:
		return "image/webp"
	case ContentTypeAutoDetect:
		fallthrough
	default:
		buff := []byte(content[:512])
		return http.DetectContentType(buff)
	}
}
