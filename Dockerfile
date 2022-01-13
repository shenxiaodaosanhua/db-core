FROM golang:1.17.5-alpine3.15 as builder
RUN mkdir "/src"
ADD . "/src"
WORKDIR "/src"
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod tidy && go build -o db-core main.go && chmod +x db-core

FROM alpine:3.15
ENV ZONEINFO=/app/zoneinfo.zip
RUN mkdir "/app"
WORKDIR "/app"
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /app
COPY --from=builder /src/db-core /app
ENTRYPOINT ["./db-core"]
EXPOSE 8081 8090
