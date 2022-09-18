FROM golang:1.17.1 as builder

WORKDIR /app

EXPOSE 80

COPY . .

RUN apt-get update && apt-get install jq -y && chmod +x ./wait
# Если что-то не собирается из-за CGO, может быть, при проверке сертификатов из гошки.
# Убрать CGO_ENABLED
# Итоговый image взять с gcc, например ubuntu
RUN CGO_ENABLED=0 make build

FROM alpine

COPY --from=builder /app /app

WORKDIR /app

CMD ls && ./wait && ./server.out
