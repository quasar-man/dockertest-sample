FROM golang:1.23.0

WORKDIR /app
COPY ./ /app

RUN go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 8080

CMD ["go", "run", "main.go"]
