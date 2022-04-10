FROM golang:1.17.1

WORKDIR /app

COPY . .

RUN apt-get clean
RUN apt-get update
RUN apt-get install jq -y

EXPOSE 8080

RUN make build

RUN chmod +x ./wait

CMD ./wait && ./server.out