package subscriber

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsMsg struct {
	q             *sqsQueue
	messageID     string
	receiptHandle string
	payload       []byte
}

func (m *sqsMsg) Payload() []byte {
	return m.payload
}

func (m *sqsMsg) ChangeVisibility(ctx context.Context, remainingInvisibility time.Duration) error {

	params := &sqs.ChangeMessageVisibilityInput{
		QueueUrl:          aws.String(m.q.url),
		ReceiptHandle:     aws.String(m.receiptHandle),
		VisibilityTimeout: aws.Int64(int64(remainingInvisibility / time.Second)),
	}
	if _, err := m.q.cli.ChangeMessageVisibility(params); err != nil {
		return err
	}
	return nil
}

func (m *sqsMsg) Delete(ctx context.Context) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(m.q.url),
		ReceiptHandle: aws.String(m.receiptHandle),
	}
	if _, err := m.q.cli.DeleteMessage(params); err != nil {
		return err
	}
	log.Println("message deleted")
	return nil
}
