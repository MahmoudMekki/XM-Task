FROM golang:latest

WORKDIR /XM-Task

COPY . .
CMD ["go", "run", "services/workers/activitylog.go"]