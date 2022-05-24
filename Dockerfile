FROM golang:1.18

RUN mkdir /data

WORKDIR /go/src/pagos

COPY . .

RUN go env -w GO111MODULE=off

RUN go get -d

RUN go build -o /go/bin/pagos

ENTRYPOINT ["/go/bin/pagos"]