FROM golang:1.17.1

WORKDIR /app

COPY . .

RUN apt-get clean
RUN apt-get update
RUN apt-get install jq -y

EXPOSE 80

RUN make build

RUN chmod +x ./wait

CMD ./wait && ./server.out

#
#FROM golang:1.17.1 as builder
#
#WORKDIR /app
#
#COPY . .
#
#RUN apt-get clean
#RUN apt-get update
#RUN apt-get install jq -y
#
#EXPOSE 80
#
#RUN make build
#
#FROM alpine
#
#COPY --from=builder /app .
#
#RUN ls -la
#
#RUN chmod +x ./wait
#
#RUN chmod +x ./server.out
#
#CMD ./wait && ./server.out
#
# Почему-то main_1             | /bin/sh: ./server.out: not found