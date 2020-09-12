package bus

import "encoding/base64"

// The SQS/SNS API uses strings to represent message payloads. This function encodes a byte slice to a string.
func encodePayload(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

// The SQS/SNS API uses strings to represent message payloads. This function decodes a message payload
// into a byte slice.
func decodePayload(msgPayload string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(msgPayload)
	if err != nil {
		return nil, err
	}
	return data, nil
}
