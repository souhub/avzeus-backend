FROM golang:1.15.0-alpine3.12

WORKDIR /go/src/github.com/souhub/avzeus-backend

COPY . .

RUN go build cmd

EXPOSE 8000

CMD ["go","run","cmd/main.go"]
