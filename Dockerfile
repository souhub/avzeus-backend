FROM golang:1.15.0-alpine3.12

WORKDIR /go/src/github.com/souhub/avzeus-backend

COPY . .

RUN go build cmd

ENV CLOUD_STORAGE_PATH https://storage.googleapis.com/avzeus

ENV BACKEND_URL http://localhost:8000

ENV AI_URL http://localhost:5000

ENV FRONTEND_URL https://avzeus-client.mmu6fa6rgrojg.ap-northeast-1.cs.amazonlightsail.com

EXPOSE 8000

CMD ["go","run","cmd/main.go"]
