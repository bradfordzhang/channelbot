FROM golang:alpine

WORKDIR /build

COPY main.go .

COPY go.mod .

COPY go.sum .

RUN go build -o zycschannelbot main.go

CMD ["./zycschannelbot"]
