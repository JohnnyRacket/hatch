FROM golang:latest
ADD . /go/src/app-service
WORKDIR /go/src/app-service
RUN go get 
RUN go build
# start app
ENTRYPOINT [ "./app-service" ]