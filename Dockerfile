FROM golang:latest

WORKDIR /go/src/ytm_search
COPY . .

RUN go get -d -v -u ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["ytm_search"]
