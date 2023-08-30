FROM ubuntu:latest

WORKDIR /app

COPY . /app

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y ca-certificates
RUN apt-get install golang-go -y
RUN go get -u github.com/cosmtrek/air
COPY ./main.sh /app/main.sh
RUN chmod +x /app/main.sh

RUN chmod +x *.go

CMD ["/app/main.sh"]
