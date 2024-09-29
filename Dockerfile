FROM golang:1.23.1 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /crud ./cmd/main.go

EXPOSE 8000

CMD [ "/crud" ]