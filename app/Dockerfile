FROM golang:1.22-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go get github.com/githubnemo/CompileDaemon

RUN go install github.com/githubnemo/CompileDaemon

CMD [ "sh", "entrypoint.sh" ]