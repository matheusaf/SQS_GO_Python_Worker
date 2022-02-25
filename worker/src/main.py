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
        sqs_queue = sqs.sqs_resource()

        while True:
            try:
                # fila
                messages = sqs_queue.receive_messages(
                    # recupera todos os atributos
                    MessageAttributeNames=['All'],
                    MaxNumberOfMessages=10,
                    WaitTimeSeconds=5,
                )

                for message in messages:
                    logger.info(f"body: { message.body }, attr: { message.message_attributes }")
                    parsed_message = sqs.retrive_message_attributes(message)

                    print(parsed_message)
                    
                    result = slack.send_message(slack_client, parsed_message.get("Channel"), parsed_message.get("Message"))
                    
                    message.delete()

                    if result is None:
                        logger.info(result)
                        continue

                    logger.info("Channel not found, message deleted")

            except sqs.ClientError as sqs_err:
                logger.error(f"SQS -> Error while using SQS: '{sqs_err}'")
                logger.error(traceback.print_exc())

            except slack.SlackApiError as slack_err:
                logger.error(f"Slack -> Error while using Slack: '{slack_err}'")
                logger.error(traceback.print_exc())
			
            except Exception as err:
                logger.error(f"Error posting message: {err}")
                logger.error(traceback.print_exc())

    except Exception as main_exc:
        logger.error(f"Error posting message: {main_exc}")

if __name__ == "__main__":
	main()
