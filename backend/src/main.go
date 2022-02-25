package main

import (
	"net/http"
	"teste_ifood/model"
	"teste_ifood/sqs"
	"teste_ifood/view"

	"github.com/gin-gonic/gin"
)

func main() {
	// inicializando o servidor
	router := gin.Default()

	// criando grupo de rotas
	v1 := router.Group("/v1")
	{
		// criando rota para receber a mensage
		// esta funcao recebe um POST request contendo a mensagem no body
		v1.POST("/message", func(c *gin.Context) {
			var message model.MessageModel

			// realizando o parse do json para o modelo de mensagem
			// caso apresente algum erro o servidor retornara 400 com a mensagem informando o erro
			if err := c.ShouldBindJSON(&message); err != nil || message.Channel == "" || message.Message == "" {
				c.JSON(http.StatusBadRequest, view.ErrorMessage{
					Message: "Mensagem recebida apresenta um formato inválido ou não apresenta valores",
					Error:   err.Error()})
				// retorna antes para evitar a execucao do bloco abaixo
				return
			}

			// Realizando o envio da mensagem para a fila e armazenando a resultado e o erro retornado
			msgOutput, err := sqs.SendMessage(message)

			// Caso tenha ocorrido algum erro ao enviar a mensagem o servidor ira retornar 500 com a mensagem de erro
			if err != nil {
				c.JSON(http.StatusInternalServerError, view.ErrorMessage{
					Message: "Erro ao enviar mensagem",
					Error:   err.Error()})
				// retorna antes para evitar a execucao do bloco abaixo
				return
			}


			// Caso o model recebido esteja correto e a mensagem seje enviada com sucesso o servidor retornara 200 com os itens da mensagem
			c.JSON(http.StatusOK, view.MessageView{
				MessageId:         *msgOutput.MessageId,
				MessageBody:       *msgOutput.MD5OfMessageBody,
				MessageAttributes: *msgOutput.MD5OfMessageAttributes})

		})
	}

	// executando o servidor na porta 80
	router.Run(":80")
}
