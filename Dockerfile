# Use an official Go runtime as a parent image
FROM golang:1.20-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app


WORKDIR /app/cmd/snixr


# Build the Go app
RUN go build -o snixr

# Expose port 3000 for the application
EXPOSE 3000

# Define the command to run when the container starts
CMD ["./snixr"]

