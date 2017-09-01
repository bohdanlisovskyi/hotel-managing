# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.8.3
# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/github.com/bohdanlisovskyi/hotel-managing
# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY . /go/src/github.com/bohdanlisovskyi/hotel-managing
WORKDIR /go/src/github.com/bohdanlisovskyi/hotel-managing
# Download and install any required third party dependencies into the container.
RUN go get -u github.com/Masterminds/glide && glide install
# Set the PORT environment variable inside the container
ENV PORT 3001

EXPOSE 3001
RUN go build -o hotel_managing
CMD ./hotel_managing
