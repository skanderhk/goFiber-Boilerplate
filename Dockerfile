# Building the binary of the App
FROM golang:1.21.5-alpine AS builder

# `boilerplate` should be replaced with your project name
WORKDIR /go/src/boilerplate

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod vendor

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server ./app/main.go


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:3.19.0 AS release

WORKDIR /app

# `boilerplate` should be replaced here as well
COPY --from=builder /go/src/boilerplate/server .
COPY --from=builder /go/src/boilerplate/app/public ./app/public


# Arguments default to local
ENV APP_ENV=local

RUN echo "APP_ENV=$APP_ENV" > .env

# Add packages
RUN apk -U upgrade \
    && chmod +x /app/server

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

ENTRYPOINT ["./server","-env=prod" ]