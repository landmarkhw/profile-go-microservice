FROM golang:latest
WORKDIR /app/src/github.com/landmarkhw/profile-go-microservice
ENV GOPATH=/app

# Copy files to the container
COPY ./ ./

# Restore dependencies
RUN go get ./...
# Build the app
RUN go build -o main

# Run the app
CMD ./main