package main

import (
	"net/http"
	"teste_ifood/model"
	"teste_ifood/sqs"
	"teste_ifood/view"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/message", func(c *gin.Context) {
			var message model.MessageModel

			if err := c.ShouldBindJSON(&message); err != nil || message.Channel == "" || message.Message == "" {
				c.JSON(http.StatusBadRequest, view.ErrorMessage{
					Message: "Mensagem recebida apresenta um formato inválido ou não apresenta valores",
					Error:   err.Error()})
				return
			}

			msgOutput, err := sqs.SendMessage(message)

			if err != nil {
				c.JSON(http.StatusInternalServerError, view.ErrorMessage{
					Message: "Erro ao enviar mensagem",
					Error:   err.Error()})
				return
			}

			c.JSON(http.StatusOK, view.MessageView{
				MessageId:         *msgOutput.MessageId,
				MessageBody:       *msgOutput.MD5OfMessageBody,
				MessageAttributes: *msgOutput.MD5OfMessageAttributes})

		})
	}

	router.Run(":80")
}
