##BUILD stage

FROM golang:1.14 as build
WORKDIR /go/src/docker-demo-go
COPY vendor .
COPY . .
RUN make build-linux

##RUN stage
FROM alpine
WORKDIR /go/src/docker-demo-go
COPY --from=build /go/src/docker-demo-go/out/docker-demo-go .
CMD ["./docker-demo-go"]