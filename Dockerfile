FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o task ./cmd/task

FROM alpine

WORKDIR /build 

COPY ./configs ./configs 
COPY --from=builder /build/task /build/task

ARG CONFIG="configs/test1.txt"
ENV CONFIG ${CONFIG}
CMD ./task --config-path ${CONFIG}