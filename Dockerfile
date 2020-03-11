FROM golang:1.14.0-alpine3.11
EXPOSE 8080
WORKDIR /go/src/app
ADD ./ ./
RUN echo $GOPATH && \
    apk add --no-cache git && \
    apk add --no-cache libxslt && \
    go build -o workers
CMD nohup ./workers >> workers.log 2>&1
