package bus

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/evoke365/webserver/store/data"
)

type SNS struct {
	cli *sns.SNS
}

func NewSNS(sess *session.Session) *SNS {
	return &SNS{
		cli: sns.New(sess),
	}
}

func (s *SNS) Publish(topic data.EventType, message []byte) error {
	input := &sns.PublishInput{
		Message:  aws.String(encodePayload(message)),
		TopicArn: aws.String(string(topic)),
	}

	if _, err := s.cli.Publish(input); err != nil {
		return err
	}

	return nil
}
