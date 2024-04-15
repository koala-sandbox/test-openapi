FROM golang:1.20-alpine

WORKDIR /app

# https://goreleaser.com/errors/docker-build/ - see Do's and Don'ts
COPY server .
COPY VERSION . 
ENTRYPOINT ["./server"]
