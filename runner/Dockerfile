FROM golang:1.10

RUN curl -fsSLO --compressed https://github.com/github/gh-ost/releases/download/v1.0.46/gh-ost-binary-linux-20180527215024.tar.gz && \
    tar -xf gh-ost-binary-linux-20180527215024.tar.gz -C /bin

ENV GHOST_PATH "/bin/gh-ost"

WORKDIR /go/src/github.com/Buzzvil/gastly/runner
ADD . /go/src/github.com/Buzzvil/gastly/runner

RUN go build ./cmd/runner
CMD ["./runner"]
EXPOSE 8000
