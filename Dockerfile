#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
COPY --from=builder /go/bin/szkuvify /szkuvify
LABEL Name=szkuvify Version=0.0.1
CMD [ "/szkuvify" ]