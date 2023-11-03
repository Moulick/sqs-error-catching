package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to get AWS config: %v", err)
	}

	stsClient := sts.NewFromConfig(cfg)
	whoami, err := stsClient.GetCallerIdentity(context.Background(), nil)
	if err != nil {
		log.Fatalf("unable to get caller identity: %v", err)
	}
	accountID := *whoami.Account

	sqsClient := sqs.NewFromConfig(cfg)

	_, err = sqsClient.DeleteQueue(context.Background(), &sqs.DeleteQueueInput{
		QueueUrl: aws.String("https://sqs.eu-central-1.amazonaws.com/" + accountID + "/random-dummy-sqs"),
	})

	var dne *types.QueueDoesNotExist
	if err != nil {
		if errors.As(err, &dne) {
			// Queue does not exist
			fmt.Println("Queue does not exist")
		}
		log.Fatalf("unable to delete queue: %v", err)
	}
}
