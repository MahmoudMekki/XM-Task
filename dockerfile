FROM golang:latest

WORKDIR /XM-Task

COPY . .

RUN go mod tidy
RUN go mod vendor


CMD ["go", "test", "./tests/company/..."]
CMD ["go", "run", "services/api/main.go"]