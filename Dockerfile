FROM arm32v7/ubuntu:latest
LABEL maintainer="iw4p@protonmail.com"

RUN apt-get update && \
    apt-get -y install gcc make && \
    apt-get -y install git && \
    apt-get -y install golang-1.18 && \
    apt-get -y install gcc-arm-linux-gnueabihf

ENV PATH="/usr/lib/go-1.18/bin:$PATH"

COPY . /home/root/
WORKDIR /home/root/

RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-linux-gnueabihf-gcc go build -o main

ENV LD_LIBRARY_PATH=./arm
RUN ./main
