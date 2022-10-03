package main

import (
	sqs "aws-sqs-lambda/types"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(_ context.Context, sqsEvent events.SQSEvent) {
	for _, message := range sqsEvent.Records {

		var messageBody sqs.Body

		if err := json.Unmarshal([]byte(message.Body), &messageBody); err != nil {
			fmt.Printf("❌ Unable to parse SQS body: %s\n", err)
			continue
		}

		for _, record := range messageBody.Records {
			s3Entity := record.S3
			fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3Entity.Bucket.Name, s3Entity.Object.Key)
		}
	}
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
