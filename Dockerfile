FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/baopham/summarizer
COPY . /usr/local/go/src/github.com/baopham/summarizer
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/summarizer ./summarizer


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/baopham/summarizer/build/summarizer /bin/summarizer
CMD ["summarizer", "up"]
