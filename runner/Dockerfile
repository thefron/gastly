FROM golang:1.10

WORKDIR /go/src/github.com/Buzzvil/gastly/runner
ADD . /go/src/github.com/Buzzvil/gastly/runner

RUN go build ./cmd/runner
CMD ["./runner"]
EXPOSE 8000
