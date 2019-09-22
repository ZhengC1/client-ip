FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o client_app ./client_ip/src/app/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./client_app"]
