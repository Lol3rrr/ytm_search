FROM golang:latest

WORKDIR /go/src/ytm_search
COPY . .

RUN go get -d ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["ytm_search"]
