# Integração SQS, GO e Python

Neste projeto foi realizado os seguintes desenvolvimentos:

- Desenvolvimento do backend em Go com a finalidade de receber um request POST contendo a mensagem e o canal ao ser enviado para a fila do SQS
- Configuração do Localstack para utilizar o SQS
- Desenvolvimento de um worker em Python para verificar e consumir a fila do SQS e em sequência enviar as mensagens presente na fila para o canal presente na mensagem recebida

Cada recurso desenvolvido possui um container

Para executar o projeto, basta utilizar o comando abaixo no diretório raiz do projeto:
'''console
docker-compose up
'''
