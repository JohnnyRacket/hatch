FROM golang:latest
ADD . /go/src/hatch/app-service
WORKDIR /go/src/hatch/app-service
RUN go get 
RUN go build
# start app
ENTRYPOINT [ "./app-service" ]