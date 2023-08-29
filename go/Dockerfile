FROM ubuntu:latest

WORKDIR /app

COPY . .

RUN apt-get update -y
RUN apt-get upgrade -y

RUN apt-get install golang-go -y

RUN chmod +x *.go

CMD [ "go", "run" ,"." ]
