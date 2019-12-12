FROM golang:1.10 AS builder

WORKDIR /go/src/github.com/next-lucasmenendez/interpretext-api

COPY . .

ENV MODELS="/go/src/github.com/next-lucasmenendez/interpretext-api/models"

RUN go get github.com/tools/godep && \
    godep get && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o interpretext . && \
    mkdir models && \
    wget https://raw.githubusercontent.com/next-lucasmenendez/interpretext-postagger/master/es && \
    wget https://raw.githubusercontent.com/next-lucasmenendez/interpretext-postagger/master/en && \
    ./interpretext train ./es ./en

FROM alpine

COPY --from=builder /go/src/github.com/next-lucasmenendez/interpretext-api/interpretext /go/bin/interpretext
COPY --from=builder /go/src/github.com/next-lucasmenendez/interpretext-api/models /go/bin/models

ENV PORT="80" \
    MODELS="/go/bin/models" \
    STOPWORDS=""

CMD ["/go/bin/interpretext", "server"]

EXPOSE 80
