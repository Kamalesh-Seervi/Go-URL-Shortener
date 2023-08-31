FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
# RUN go get -u github.com/cosmtrek/air

COPY . .

# RUN go build -o main .

EXPOSE 7777
# ENTRYPOINT ["air"]

CMD ["go" ,"run" ,"."]