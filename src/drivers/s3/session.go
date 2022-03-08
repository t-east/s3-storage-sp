package s3

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Session() *session.Session {
	// sessionの作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}

func UploadContentS3(key string, file []byte) error {
	// sessionの作成
	// TODO: drivers層のsessionを渡す
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))
	f := bytes.NewReader(file)

	bucketName := "xxx-bucket"
	objectKey := key

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   f,
	})
	if err != nil {
		return err
	}
	return nil
}

func DownLoadContentS3(key string) (*s3.GetObjectOutput, error) {
	// sessionの作成
	// TODO: drivers層のsessionを渡す
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))

	bucketName := "xxx-bucket"
	objectKey := key

	// S3 clientを作成
	svc := s3.New(sess)

	// S3からファイルをダウンロードせずに読み込む
	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				err := fmt.Sprintf("bucket %s does not exist", bucketName)
				return nil, errors.New(err)
			case s3.ErrCodeNoSuchKey:
				err := fmt.Sprintf("object with key %s does not exist in bucket %s", objectKey, bucketName)
				return nil, errors.New(err)
			default:
				return nil, aerr
			}
		}
	}
	return obj, nil
}
