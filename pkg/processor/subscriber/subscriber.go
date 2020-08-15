package subscriber

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Client interface {
	ReceiveMessage(ctx context.Context) (Message, error)
}

// Message defines the beavious of a message of an order queue.
type Message interface {
	Payload() []byte

	// Resets the remaining visibility timeout to the given value.
	ChangeVisibility(ctx context.Context, remainingInvisibility time.Duration) error

	// Deletes the message. Once this is called, no further methods may be called
	// on this message.
	Delete(ctx context.Context) error
}

type SQSClient struct {
	client *sqs.SQS
}

func NewSQSClient(awsSess *session.Session) *SQSClient {
	return &SQSClient{
		client: sqs.New(awsSess),
	}
}
