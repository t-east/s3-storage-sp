package s3
import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func Session() *session.Session {
	// sessionの作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}