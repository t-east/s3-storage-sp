package storage

import (
	"sp/src/domains/entities"
	"sp/src/usecases/port"

	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type ContentStorage struct {
}

func NewContentStorage() port.ContentStorage {
	return &ContentStorage{}
}

func (pr *ContentStorage) UploadContentS3(content *entities.Content) error {
	// sessionの作成
	// TODO: drivers層のsessionを渡す
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))
	file := bytes.NewReader(content.Content)

	bucketName := "xxx-bucket"
	objectKey := content.ContentId + "/" + content.ContentName

	// Uploaderを作成し、ローカルファイルをアップロード
	// TODO: Upload関数もdrivers層に格納する
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return err
	}

	for i := 0; i < content.SplitCount; i++ {
		metaKey := content.ContentId + "/" + content.ContentName + "/" + fmt.Sprint(i)
		uploader = s3manager.NewUploader(sess)
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(metaKey),
			Body:   file,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (pr *ContentStorage) Create(c *entities.Content) (*entities.Content, error) {
	WriteBinaryFile(c.Id, binary.BigEndian, c.Content)
	return c, nil
}

func (pr *ContentStorage) Get(id string) (*entities.Content, error) {
	return &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          id,
		UserId:      id,
		ContentId:   id,
	}, nil
}

/*
バイナリデータに変換してファイルに出力する
*/
func WriteBinaryFile(filename string, order binary.ByteOrder, val interface{}) {
	// バイナリデータの格納用
	buf := new(bytes.Buffer)
	// valの値をバイナリデータに変換してbufに格納する
	err := binary.Write(buf, order, val)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// ファイル作成
	file, err2 := os.Create(filename)
	if err2 != nil {
		fmt.Println("file create err:", err2)
		return
	}
	// バイナリデータをファイルに書き込み
	_, err3 := file.Write(buf.Bytes())
	if err3 != nil {
		fmt.Println("file write err:", err3)
		return
	}
	fmt.Println("file write ok.")
}
