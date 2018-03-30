FROM golang
WORKDIR /go/src/github.com/zoidbergwill/docker-counter
RUN go get -d -v gopkg.in/redis.v5
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/zoidbergwill/docker-counter/app .
ENV DOCKER_COUNTER_ADDR 0.0.0.0
EXPOSE 8080
CMD ["./app"]
