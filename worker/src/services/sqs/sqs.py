#!/usr/bin/env python3

import os
import queue
from typing import Dict

import boto3
from botocore.exceptions import ClientError


def retrive_message_attributes(queue_message) -> Dict[str, str]:
    """
    Send a message to an Amazon SQS queue.

    :param queue: The queue that receives the message.
    :param message_body: The body text of the message.
    :param message_attributes: Custom attributes of the message. These are key-value pairs that can be whatever you want.
    :return: The response from SQS that contains the assigned message ID.
    """
    message: Dict[str, Dict[str, str]] = queue_message.message_attributes

    if not message:
        return {}

    return { 
        key: value.get("StringValue", "") 
        for key, value in message.items()
    }


def sqs_resource() -> object:
    # iniciando a fila
    aws_region = os.environ.get("AWS_DEFAULT_REGION")
    endpoint_url = os.environ.get("AWS_ENDPOINT")
    aws_access_key_id = os.environ.get("AWS_ACCESS_KEY_ID")
    aws_secret_access_key = os.environ.get("AWS_SECRET_ACCESS_KEY")
    queue_name: str = os.environ.get("SQS_QUEUE_NAME")

    # configurando o long pooling do sqs
    sqs = boto3.resource(
        "sqs", 
		region_name=aws_region, 
        endpoint_url=endpoint_url,
        aws_access_key_id=aws_access_key_id,
        aws_secret_access_key=aws_secret_access_key,
    )
	
    # recuperando a fila
    return sqs.get_queue_by_name(QueueName=queue_name)
