FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/snipbox

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
# RUN CGO_ENABLED=0 go test -v ./...

# Build the Go app
RUN go build -o ./out/snipbox /tmp/snipbox/cmd/web

# Start fresh from a smaller image
FROM alpine:3.9
ARG PORT=4000

COPY --from=builder /tmp/snipbox/out/snipbox /app/snipbox

# This container exposes port 4000 to the outside world
EXPOSE $PORT

# Run the binary program produced by `go install`
CMD ["/app/snipbox"]
