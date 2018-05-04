FROM golang:latest
ADD . /go/src/hatch/user-service
WORKDIR /go/src/hatch/user-service
RUN go get 
RUN go build
ENTRYPOINT [ "./user-service" ]