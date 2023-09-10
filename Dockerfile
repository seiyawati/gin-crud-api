FROM golang:latest

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apt-get update && apt-get install -y git
RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
