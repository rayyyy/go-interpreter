FROM golang:1.13.4

RUN apt-get update && apt-get install -y --no-install-recommends zsh
RUN go get golang.org/x/tools/cmd/godoc

RUN mkdir /app
WORKDIR /app
