FROM golang:1.18.6

WORKDIR /app

ADD . /app

#CMD "./mailganer"

EXPOSE 8080 8080