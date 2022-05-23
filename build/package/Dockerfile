# Dockerfile for:
# - project: packet-monitoring
# - maintainer: amirhnajafiz
# - date: 2022

# form alpine golang 1.17 start our image building
FROM golang:1.17-alpine

# app workdirectory
WORKDIR /app

# copy every file
COPY . .

# building our application
RUN go build -o /main

# expose on port 1224
EXPOSE 1224

# enter container by executing the main process
ENTRYPOINT ["/main"]