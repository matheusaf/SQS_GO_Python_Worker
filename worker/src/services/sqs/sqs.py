#!/usr/bin/env python3

import os
import queue
from typing import Dict

import boto3
from botocore.exceptions import ClientError


def retrive_message_attributes(queue_message) -> Dict[str, str]:
    # recuperando atributos da mensagem
    message: Dict[str, Dict[str, str]] = queue_message.message_attributes

    # caso a mensagem seja nula é retornado um dicionario vazio
    if not message:
        return {}

    # retornando um dicionario contendo apenas os nomes dos atributos e seus respectivos valores
    return { 
        key: value.get("StringValue", "") 
        for key, value in message.items()
    }


def sqs_resource() -> object:
    # obtendo variaveis de ambiente necessarias para conectar com o SQS
    aws_region = os.environ.get("AWS_DEFAULT_REGION")
    endpoint_url = os.environ.get("AWS_ENDPOINT")
    aws_access_key_id = os.environ.get("AWS_ACCESS_KEY_ID")
    aws_secret_access_key = os.environ.get("AWS_SECRET_ACCESS_KEY")
    queue_name: str = os.environ.get("SQS_QUEUE_NAME")

    # criando conexao com o locastack utilizando as variaveis de ambiente obtidas acima
    sqs = boto3.resource(
        "sqs", 
		region_name=aws_region, 
        endpoint_url=endpoint_url,
        aws_access_key_id=aws_access_key_id,
        aws_secret_access_key=aws_secret_access_key,
    )
	
    # recuperando a fila
    return sqs.get_queue_by_name(QueueName=queue_name)
