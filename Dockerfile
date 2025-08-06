FROM golang:alpine AS builder
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .
#COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.12
# Install Bash and curl
RUN apk update \
 && apk add --no-cache curl bash \
 && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app/main .
# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080
#run
CMD ["./main"]