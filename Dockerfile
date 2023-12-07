FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.10

RUN adduser -DH ningyotsukai

WORKDIR /app

COPY --from=builder /app/main /app/

COPY etc/config.yaml /app/etc/config.yaml
COPY fonts /app/fonts
RUN chown ningyotsukai:ningyotsukai /app
RUN chmod +x /app

USER ningyotsukai

CMD ["/app/main"]