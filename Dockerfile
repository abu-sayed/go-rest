FROM golang:1.14-alpine

RUN mkdir /app 
ADD . /app/
WORKDIR /app 

RUN go mod vendor
RUN go build -o main .

EXPOSE 8080
CMD ["/app/main"]