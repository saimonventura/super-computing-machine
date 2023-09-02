# Use the official Golang image as a base
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files, if they exist, into the work directory
COPY go.* ./

# Download all dependencies (this can be beneficial for caching)
RUN go mod download

# Copy the source code into the working directory
COPY . .

# Expose port 8080 (or whatever port your application uses)
EXPOSE 8080

# Comando para executar a aplicação
CMD ["go", "run", "main.go"]
