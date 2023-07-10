FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o task ./cmd/task

FROM alpine

WORKDIR /build 

COPY ./configs ./configs 
COPY --from=builder /build/task /build/task

ARG CONFIG="configs/test_1.txt"
ENV CONFIG configs/${CONFIG}

CMD ./task ${CONFIG}