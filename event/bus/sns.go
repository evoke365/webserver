package bus

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSstruct {
	cli *sns.SNS
}

func NewSNS(sess *session.Session) *SNS {
	return &SNS{
		cli: sns.New(sess),
	}
}

func (s *SNS) Publish(topic data.EventType, message []byte) error {
	input := &sns.PublishInput{
		Message: aws.String(encodePayload(payload)),
		TopicArn: aws.String(topic),
	}

	if _, err := svc.Publish(input); err != nil {
		return err
	}
	
	return nil 
}