#!/usr/bin/env python3

import logging
import traceback

from services import slack, sqs

# logger
logger: logging.Logger = logging.getLogger(__name__)

def main():
    print("Starting worker...")
    try:
        # iniciando slack
        slack_client = slack.create_slack_client()
        # iniciando SQS
        sqs_queue = sqs.sqs_resource()

        while True:
            try:
                # lendo a fila
                messages = sqs_queue.receive_messages(
                    # recupera todos os atributos
                    MessageAttributeNames=['All'],
                    MaxNumberOfMessages=10,
                    WaitTimeSeconds=5,
                )

                # percorrendo a lista de mensagens presentes na fila
                for message in messages:
                    # realizando log da mensagem recebida
                    logger.info(f"body: { message.body }, attr: { message.message_attributes }")

                    # recuperando atributos da mensagem
                    parsed_message = sqs.retrive_message_attributes(message)

                    # enviando mensagem para o slack
                    result = slack.send_message(slack_client, parsed_message.get("Channel"), parsed_message.get("Message"))
                    
                    # apagando a mensagem apos a leitura e envio da mesma
                    message.delete()

                    if result is None:
                        logger.info(result)
                        continue
                    
                    # realizando log caso o canal nao seja encontrado
                    logger.info("Channel not found, message deleted")

            # realizando log especifico caso ocorra algum erro durante a leitura da fila
            except sqs.ClientError as sqs_err:
                logger.error(f"SQS -> Error while using SQS: '{sqs_err}'")
                logger.error(traceback.print_exc())

            # realizando log especifico caso ocorra algum erro durante o envio da mensagem
            except slack.SlackApiError as slack_err:
                logger.error(f"Slack -> Error while using Slack: '{slack_err}'")
                logger.error(traceback.print_exc())
			
            # realizando log caso ocorra qualquer outro tipo de erro
            except Exception as err:
                logger.error(f"Error posting message: {err}")
                logger.error(traceback.print_exc())

    # realizando log caso ocorra algum erro ao iniciar a mensagem
    except Exception as main_exc:
        logger.error(f"Error posting message: {main_exc}")

if __name__ == "__main__":
	main()
