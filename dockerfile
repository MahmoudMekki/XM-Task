FROM golang:latest

WORKDIR /XM-Task

COPY . .

CMD ["go", "test", "./tests/company/..."]
CMD ["go", "run", "services/api/main.go"]