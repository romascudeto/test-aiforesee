FROM golang:1.10.5-alpine3.8

 

WORKDIR /home/alvin/Documents/dev-box-golang/aiforesee-boilerplate

COPY server.go .

RUN go build -o main .

EXPOSE 8000

CMD [“./main”]

