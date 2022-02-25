package sqs

import (
	"context"
	"os"

	"teste_ifood/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// Funcao responsavel por criar a comunicacao com o cliente do SQS
func GetClient(endppoint string, signingRegion string) (*sqs.Client, error) {

	// Funcao para alterar o endpoint utilizado pelo AWS CLI
	endPointResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if endppoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           endppoint,
				SigningRegion: signingRegion,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	// Carrega as configuracoes utilizando a funcao de configuracao definida anteriormente
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(endPointResolver))

	// Caso ocorra algum erro, retorna o erro e o cliente retornara nil
	if err != nil{
		return nil, err
	}

	// Cria e retorna o cliente do SQS
	return sqs.NewFromConfig(cfg), err
}

// Funcao resposavel por obter a url da fila
func GetQueue(queueName string, client *sqs.Client) (*string, error) {

	// Cria o objeto para pesquisar a fila utilizando o nome da fila
	getQueueUrl := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}

	// Funcao responsavel por obter a url da fila utilizando o objeto criado anteriormente
	result, err := client.GetQueueUrl(context.TODO(), getQueueUrl)

	// Caso ocorra algum erro, retorna o erro e a url retornara nil
	if err != nil{
		return nil, err
	}

	// Retorna a url da fila
	return result.QueueUrl, err
}

// Funcao responsavel por enviar a mensagem
func SendMessageHelper(message model.MessageModel, client *sqs.Client, queueUrl string) (*sqs.SendMessageOutput, error) {

	// Criando o objeto que representa a mensagem a ser enviada
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

	// Funcao que envia a mensagem para a fila
	resp, err := client.SendMessage(context.TODO(), sqsMessage)

	// Caso ocorra algum erro, retorna o erro e a mesangem retornara nil
	if err != nil{
		return nil, err
	}

	// Retorna a resposta do SQS apos receber a mensagem
	return resp, err
}


// Funcao consolidada que ira realizar o envio da mensagem utilizando 
// de apoio todas as funcoes definidas anteriormente
func SendMessage(message model.MessageModel) (*sqs.SendMessageOutput, error) {

	// Obtem as variaveis de ambiente necessarias para configurar o cliente do SQS
	AWS_ENDPOINT := os.Getenv("AWS_ENDPOINT")
	SQS_QUEUE_NAME := os.Getenv("SQS_QUEUE_NAME")
	AWS_DEFAULT_REGION := os.Getenv("AWS_DEFAULT_REGION")

	// Chamando a funcao definida anteriormente para criar e obter o cliente do SQS
	client, errorClient := GetClient(AWS_ENDPOINT, AWS_DEFAULT_REGION)

	// Retorna antes caso tenha obtido algum erro ao criar o cliente
	if errorClient != nil {
		return nil, errorClient
	}
	
	// Chamando a funcao definida anteriormente para obter a url da fila
	queue_url, errorUrl := GetQueue(SQS_QUEUE_NAME, client)

	// Retorna antes caso tenha obtido algum erro ao obter a url da fila
	if errorUrl != nil {
		return nil, errorUrl
	}
	
	// Retorna o resultado da funcao de apoio para enviar a mensagem
	return SendMessageHelper(message, client, *queue_url)
}
