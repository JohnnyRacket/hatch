FROM golang:latest
ADD . /go/src/hatchery
WORKDIR /go/src/hatchery
RUN go get 
RUN go build
ENTRYPOINT [ "./hatchery" ]