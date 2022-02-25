#!/usr/bin/env python3

import os
from typing import Any, Union

from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError


def create_slack_client() -> WebClient:
    #  obtendo variavel de ambiente contendo token do slack
	slack_token: str = os.environ.get("SLACK_BOT_TOKEN")
    # inicializando e retornando webclient
	return WebClient(token=slack_token)


def send_message(slack_client: WebClient, channel_name: str, message_text: str) -> Union[Any, None]:
	# procura se o nome do canal existe
    for channels in slack_client.conversations_list():
        for channel in channels["channels"]:
			# se existir envia a mensagem
            if channel.get("name") == channel_name:
                return slack_client.chat_postMessage(channel=channel["id"], text=message_text)
    # caso nao exista o canal retorna None
    return None
