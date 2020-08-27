# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Hunter Hartline <hunterhartline@gmail.com>"

# Install git to fetch dependencies
RUN apk update && apk add --no-cache git

# Set the current working directory inside container
WORKDIR /app

# Copy go mode and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to directory of container
COPY . .

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary file from prev stage and env file
COPY --from=builder /app/main . 
COPY --from=builder /app/.env .

# Expose port 8080 
EXPOSE 8080

# Run the executable
CMD ["./main"]

