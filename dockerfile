FROM golang:latest

WORKDIR /go/src/github.com/MahmoudMekki/XM-Task

COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go build -o XMTask

CMD ["go", "test", "./tests/company/..."]
CMD ["go", "run", "services/api/main.go"]
