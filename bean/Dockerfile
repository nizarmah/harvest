# use official go image
FROM golang:1.21.5-alpine

# set working dir
WORKDIR /go/src

# copy monorepo deps
COPY go.mod go.sum ./

# copy bean code
COPY bean/ ./bean/

# get deps
RUN go mod download

# get air for live reload
RUN go install github.com/cosmtrek/air@latest

# run bean using air
CMD air -c ./bean/config/.air.toml
