package sqs

import (
	"context"
	"fmt"
	"os"

	"teste_ifood/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func GetClient(endppoint string, signingRegion string) (*sqs.Client, error) {

	endPointResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if endppoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           endppoint,
				SigningRegion: signingRegion,
			}, nil
		}
		fmt.Println("Endpoint: ", endppoint)
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"), config.WithEndpointResolverWithOptions(endPointResolver))

	if err != nil{
		return nil, err
	}

	return sqs.NewFromConfig(cfg), err
}

func GetQueue(queueName string, client *sqs.Client) (*string, error) {

	getQueueUrl := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}

	result, err := client.GetQueueUrl(context.TODO(), getQueueUrl)

	if err != nil{
		return nil, err
	}

	return result.QueueUrl, err
}

func SendMessageHelper(message model.MessageModel, client *sqs.Client, queueUrl string) (*sqs.SendMessageOutput, error) {

	sqsMessage := &sqs.SendMessageInput{
		DelaySeconds: 10,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"Channel": {
				DataType:    aws.String("String"),
				StringValue: aws.String(message.Channel),
			},
			"Message": {
				DataType:    aws.String("String"),
				StringValue: aws.String(message.Message),
			},
		}, MessageBody: aws.String("Messagem enviada pelo Go"),
		QueueUrl: aws.String(queueUrl),
	}

	resp, err := client.SendMessage(context.TODO(), sqsMessage)

	if err != nil{
		return nil, err
	}

	return resp, err
}

func SendMessage(message model.MessageModel) (*sqs.SendMessageOutput, error) {

	AWS_ENDPOINT := os.Getenv("AWS_ENDPOINT")
	SQS_QUEUE_NAME := os.Getenv("SQS_QUEUE_NAME")
	AWS_DEFAULT_REGION := os.Getenv("AWS_DEFAULT_REGION")

	client, errorClient := GetClient(AWS_ENDPOINT, AWS_DEFAULT_REGION)

	if errorClient != nil {
		return nil, errorClient
	}

	queue_url, errorUrl := GetQueue(SQS_QUEUE_NAME, client)
	
	if errorUrl != nil {
		return nil, errorUrl
	}
	
		fmt.Println(*queue_url)

	return SendMessageHelper(message, client, *queue_url)
}
