FROM golang:latest
ENV APPDIR=src
WORKDIR ${APPDIR}
COPY src .
RUN go mod download
EXPOSE 80 443
CMD [ "go", "run", "main.go" ]