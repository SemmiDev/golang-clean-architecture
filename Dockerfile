FROM golang:1.18.3-alpine3.16 AS builder

LABEL maintainer="sammidev4@gmail.com"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .    

EXPOSE 3030
CMD ["./main"]