#!/bin/bash

echo awslocal sqs create-queue --queue-name $SQS_QUEUE_NAME
awslocal sqs create-queue --queue-name $SQS_QUEUE_NAME

echo awslocal sqs list-queues