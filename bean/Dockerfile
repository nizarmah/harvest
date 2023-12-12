# builder: official Go image
FROM golang:1.21.5 as builder

# builder: working directory
WORKDIR /go/src/app

# builder: copy source code
COPY go.mod go.sum ./
COPY bean/ ./bean/

# builder: download dependencies
RUN go mod download

# builder: build server
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bean/server ./bean/cmd/server/main.go

# runner: small Alpine image
FROM alpine:latest as runner

# runner: working directory
WORKDIR /bean

# runner: copy static files
COPY --from=builder /go/src/app/bean/static/ ./static/

# runner: copy built server
COPY --from=builder /go/src/app/bean/server .

# runner: run server
ENTRYPOINT [ "./server" ]
