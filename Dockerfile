FROM golang:1.15.0-alpine3.12

WORKDIR /go/src/github.com/souhub/avzeus-backend

COPY . .

RUN go build cmd

ENV DB_USER dbmasteruser

ENV DB_PASS )m&57>5XfFu]7922z3>,op1{HXoPcoX7

ENV DB_PROTOCOL tcp

ENV DB_ENDPOINT ls-4eec5a0b62700bee1c4159caae9f1ffd3b1dc668.crqfsr0qwzrh.us-east-1.rds.amazonaws.com

ENV DB_PORT 3306

ENV DB_NAME dbavzeus

ENV CLOUD_STORAGE_PATH https://storage.googleapis.com/avzeus

ENV BACKEND_URL http://localhost:8000

ENV AI_URL http://localhost:5000

ENV FRONTEND_URL https://avzeus-client.mmu6fa6rgrojg.ap-northeast-1.cs.amazonlightsail.com

EXPOSE 8000

CMD ["go","run","cmd/main.go"]
