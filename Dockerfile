FROM golang:1.18

ENV TZ="America/New_York"

WORKDIR /app

ADD . .

RUN go mod tidy
RUN go mod download

EXPOSE 8084

RUN go build -o /entrypoint

CMD [ "/entrypoint" ]
