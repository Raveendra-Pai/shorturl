FROM golang:1.16-alpine AS shorturl
RUN mkdir -p ./bin  && mkdir app
WORKDIR /app
COPY . /app/
RUN mkdir -p ./bin &&  CGO_ENABLED=0 GOOS=linux go build -o ./bin/shorturl ./src

FROM alpine:latest
WORKDIR /app
RUN mkdir -p bin && mkdir -p config
COPY --from=shorturl /app/bin/* /app/bin/
COPY --from=shorturl /app/config/* /app/config/
COPY --from=shorturl /app/logs/* /app/logs/
EXPOSE 8082
CMD ["./bin/shorturl"]
