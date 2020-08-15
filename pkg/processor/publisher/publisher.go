package publisher

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type Client interface {
	Publish(ctx context.Context, topic, msg string) error
}

type SNSClient struct {
	client *sns.SNS
}

// New returns a new shared instance of SNS Client.
func New(awsSess *session.Session) *SNSClient {
	return &SNSClient{
		client: sns.New(awsSess),
	}
}

// Publish publishes a message to a topic on SNS.
func (p *SNSClient) Publish(ctx context.Context, topic, msg string) error {
	input := &sns.PublishInput{
		Message:  aws.String(msg),
		TopicArn: aws.String(topic),
	}
	if _, err := p.client.Publish(input); err != nil {
		return err
	}
	return nil
}
