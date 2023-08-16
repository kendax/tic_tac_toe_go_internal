# Using go 1.20 as the base image
FROM golang:1.20

# Declaring the working directory inside the container
WORKDIR /usr/src/app

# Pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copying all project files to the container's /usr/src/app directory 
COPY . /usr/src/app
RUN go build 

# Command to run the application
CMD ["./app"]