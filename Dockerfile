FROM golang:latest 

EXPOSE 80

WORKDIR /usr/src/app
COPY ./ ./

RUN go build cmd/main.go 

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
