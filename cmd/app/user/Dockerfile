FROM golang:1.23-alpine AS build

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/serve cmd/app/user/*.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /bin/serve .

# Copy the user.config file
COPY --from=build /app/file ./file/

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./serve"]