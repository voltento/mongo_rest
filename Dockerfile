FROM golang

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY app app
COPY go.mod .
COPY go.sum .

RUN go mod edit -replace github.com/voltento/mongo_rest/app/config=/app
RUN go build -o service app/app.go

FROM alpine:latest
WORKDIR /app
COPY --from=0 /build/service app

CMD ["/app/app"]