FROM golang:1.14

WORKDIR /go/src/docker-demo-go
COPY . .
RUN make build
EXPOSE 8080

CMD ["./out/docker-demo-go"]