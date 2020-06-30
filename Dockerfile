# Ambiente de Producao
FROM golang:1.14 as base
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o /go/bin/my-app .
ENTRYPOINT ["/go/bin/my-app"]

# Ambiente de Desenvolvimento
FROM golang:1.14 as dev
WORKDIR /app
COPY . /app
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main


