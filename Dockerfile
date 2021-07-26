FROM docker.io/library/golang:1.13 as build
WORKDIR /work
COPY . /work/

ENV TEST_DIRECTORY="tests/"

RUN go mod download
