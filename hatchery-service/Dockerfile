FROM golang:latest
ADD . /go/src/hatch/hatchery-service
WORKDIR /go/src/hatch/hatchery-service
RUN go get 
RUN go build
ENTRYPOINT [ "./hatchery-service" ]