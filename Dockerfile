# base image
FROM golang:1.22.4

# source code directory
WORKDIR /app

# copy source code into image directory
COPY go.mod .
COPY main.go .

# install dependencies
RUN go mod download

# build source code
RUN go build -o /rate-my-media

# expose port
EXPOSE 8080

# run server
CMD [ "/rate-my-media --migrate" ]