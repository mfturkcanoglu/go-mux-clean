FROM golang:latest

LABEL maintainer="mfturkcan"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=8080

RUN go build

CMD ["./go-mux-clean"]