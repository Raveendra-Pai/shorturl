FROM golang:1.16-alpine AS shorturl
RUN mkdir -p ./bin  && mkdir app
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/shorturl ./src
EXPOSE 8082
CMD ["./bin/shorturl"]
