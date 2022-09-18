FROM golang:1.17.1 as builder

WORKDIR /app

EXPOSE 8080 443 80 5001

COPY . .

RUN CGO_ENABLED=0 make build-sessions

FROM alpine

WORKDIR /app

COPY --from=builder /app/sessions.out ./
COPY --from=builder /app/configs ./app/configs
COPY --from=builder /app/logs ./app/logs


CMD ./sessions.out
