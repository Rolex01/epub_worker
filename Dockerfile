FROM golang:1.14.0-alpine3.11
EXPOSE 80
WORKDIR /go/src/app
ADD ./ ./
RUN echo $GOPATH && \
    apk add --no-cache git && \
    go build -o workers
CMD nohup ./workers >> workers.log 2>&1
