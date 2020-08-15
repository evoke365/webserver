package subscriber

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// An SQS backed queue implementation.
type sqsQueue struct {
	// The AWS SQK SQS client.
	cli *sqs.SQS
	// The queue url.
	url string
	// The maximum time to wait on ReceiveMessage.
	waitTime time.Duration
}

// Config defines dependent variables and values for initialising a sqsQueue client.
type Config struct {
	// A valid aws region string must be provided to initialise a sqs client.
	// The available regions are:
	// https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html
	Region string
	// the endpoint of the queue.
	Endpoint string
	// wait time is important for reducing cost by enabling long pooling.
	// https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/reducing-costs.html
	WaitTime time.Duration
	// Using default aws credentials for live environment.
	// Create a static credentials for local development.
	// https://docs.aws.amazon.com/sdk-for-go/api/aws/credentials/
	Credentials *credentials.Credentials
}

// NewSqsQueue creates a new SQS queue instance bound to the specified url. waitTime is the number of seconds
// that an sqsl.ReceiveMessage() should wait, at most, for a message to arrive. If it is set to a
// non-zero number then long-polling is enabled, as described here:
// http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html
func NewSqsQueue(sqsCli *sqs.SQS, c Config) *sqsQueue {
	return &sqsQueue{
		cli:      sqsCli,
		url:      c.Endpoint,
		waitTime: c.WaitTime,
	}
}

func (s *sqsQueue) ReceiveMessage(ctx context.Context) (Message, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(s.url),
		MaxNumberOfMessages: aws.Int64(1),
		WaitTimeSeconds:     aws.Int64(int64(s.waitTime / time.Second)),
	}

	// Poll AWS for a message. Each call to the ReceiveMessage endpoint will block for at most
	// s.waitTime.
	awsMsg, err := func() (*sqs.Message, error) {
		for i := 1; true; i++ {
			resp, err := s.cli.ReceiveMessage(params)
			if err != nil {
				return nil, err
			}
			if len(resp.Messages) == 1 {
				return resp.Messages[0], nil
			}
			if len(resp.Messages) > 1 {
				return nil, fmt.Errorf("Error: Got %d messages from SQS, expected at most 1", len(resp.Messages))
			}
		}
		return nil, errors.New("unreachable statement")
	}()
	if err != nil {
		return nil, err
	}

	payload, err := decodePayload(*awsMsg.Body)
	if err != nil {
		return nil, err
	}

	msg := &sqsMsg{
		q:             s,
		messageID:     *awsMsg.MessageId,
		receiptHandle: *awsMsg.ReceiptHandle,
		payload:       payload,
	}
	return msg, nil
}

// The SQS API uses strings to represent message payloads. This function decodes a message payload
// into a byte slice.
func decodePayload(msgPayload string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(msgPayload)
	if err != nil {
		return nil, err
	}
	return data, nil
}
