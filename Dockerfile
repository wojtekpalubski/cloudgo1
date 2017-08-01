FROM golang:1.8.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git
RUN go get github.com/gin-gonic/gin

ENV SOURCES /go/cloudgo1
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build
WORKDIR ${SOURCES}
CMD ${SOURCES}/cloudgo1
EXPOSE 8080

