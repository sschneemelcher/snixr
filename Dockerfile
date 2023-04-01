# Use an official Go runtime as a parent image
FROM golang:1.20-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build cmd/snixr/snixr.go

# Expose port 80 for the application
EXPOSE 80

# Define the command to run when the container starts
CMD ["./snixr"]

