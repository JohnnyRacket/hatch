FROM golang:latest
ADD ./user-service /go/src/hatch/user-service
ADD ./rpc /go/src/hatch/rpc
WORKDIR /go/src/hatch/user-service
RUN go get 
RUN go build
ENTRYPOINT [ "./user-service" ]