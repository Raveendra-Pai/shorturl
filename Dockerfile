FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
ADD src ./src
ADD config ./config
RUN mkdir -p ./bin

RUN go build -o ./bin/shorturl ./src

EXPOSE 8082

CMD ["./bin/shorturl"]



