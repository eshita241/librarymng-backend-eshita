# Use the official Golang image as builder stage
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download module dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Install dependencies
RUN go get

# Build the Go binary with CGO disabled for static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o librarymng-backend

# Start a new stage
FROM golang:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage to the current working directory
COPY --from=builder /app/librarymng-backend . 

# Copy the .env file from the builder stage to the current working directory
COPY --from=builder /app/.env .

# Expose port 8080 (this is just a declaration and doesn't actually publish the port)
EXPOSE 8080

# Set the default command to run when the container starts
CMD ["./librarymng-backend"]
