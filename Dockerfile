FROM golang:1.20.5-alpine3.17 as builder
WORKDIR /app
COPY . . 
RUN go build -o main main.go

EXPOSE 8080
CMD [ "/app/main" ]


# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]
