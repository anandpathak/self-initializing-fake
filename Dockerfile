FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o self-initializing-fake .

EXPOSE 8112
EXPOSE 8113

CMD ["./self-initializing-fake"]
