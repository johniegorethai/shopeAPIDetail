# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Johnie Gorethai <jonigo07@gmail.com>"

RUN mkdir /go/src/shopeAPIDetail

WORKDIR /go/src/shopeAPIDetail

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies.
RUN go mod verify

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/cron/main.go


# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:latest
ENV ENV="staging"
ENV VERSION="1.0.1"
ENV TZ="Asia/Jakarta"
RUN apk add --no-cache tzdata
COPY --from=builder /go/src/shopeAPIDetailh/main /
COPY --from=builder /go/src/shopeAPIDetail/files/etc/orderPelapak/orderPelapak.staging.yaml /
EXPOSE 4444
ENTRYPOINT /main