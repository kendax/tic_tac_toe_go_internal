# Using go 1.20 as the base image
FROM golang:1.20

# Declaring the working directory inside the container
WORKDIR /usr/src/app

# Copying all project files to the container's /usr/src/app directory 
COPY . /usr/src/app

# Running the program's build commands and outputting (-o) the result of 'go build' to 'app'
RUN go mod download && go mod verify && go build -o app

# Command to run the program
CMD ["./app"]