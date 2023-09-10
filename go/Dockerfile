FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
# RUN go get -u github.com/cosmtrek/air

COPY . .

# RUN go build -o main .

ENV PORT=9000

# Expose the port on which the application will listen
EXPOSE $PORT
# ENTRYPOINT ["air"]

CMD ["go" ,"run" ,"."]