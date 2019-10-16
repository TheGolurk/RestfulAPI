FROM golang:latest

WORKDIR /go/src/github.com/TheGolurk/RestfulAPI
COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD [ "RestfulAPI" ]