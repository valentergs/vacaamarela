# FROM golang:1.12
# WORKDIR /go/src/app
# COPY /main .

# RUN go get -d -v ./...
# RUN go install -v ./...

# CMD ["/main"]

FROM golang:alpine

ADD main /

CMD ["/main"]
