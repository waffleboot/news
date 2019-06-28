
FROM golang

WORKDIR /go/src/github.com/waffleboot/news
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["/go/bin/storage"]
