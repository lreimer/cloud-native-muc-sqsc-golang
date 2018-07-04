[![Build Status](https://travis-ci.org/lreimer/cloud-native-muc-sqsc-golang.svg?branch=master)](https://travis-ci.org/lreimer/cloud-native-muc-sqsc-golang)

# Enterprise Cloud Native for SME - Go Service

Simple Go demo service for the Cloud Native Night with SquareScale.

https://www.meetup.com/de-DE/cloud-native-muc/

## Building and Running

```
$ go get github.com/gin-gonic/gin
$ go build -o app
$ ./app
```

## Containerizing

```
$ CGO_ENABLED=0 GOOS=linux go build -o app
$ docker build -t cloud-native-muc-sqsc-golang:1.1 .
$ docker run -it -p 9090:9090 cloud-native-muc-sqsc-golang:1.1
```
