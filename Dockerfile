FROM golang:alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk add --no-cache git && \
    go install github.com/cosmtrek/air@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
