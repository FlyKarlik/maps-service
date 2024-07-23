FROM golang:1.19.6-alpine3.16 AS builder

RUN mkdir /app
ADD ./maps-service /app
ADD ./comet /comet
ADD ./protos /protos
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine AS production
COPY --from=builder /app/main .

RUN mkdir /cert
COPY --from=builder /app/cert/service.pem /cert/service.pem
COPY --from=builder /app/cert/service.key /cert/service.key
COPY --from=builder /app/cert/ca.cert /cert/ca.cert

EXPOSE 3000
CMD ["./main"]