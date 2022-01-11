FROM golang:1.17-alpine

WORKDIR /app

RUN mkdir goAPI

COPY goAPI/go.mod ./
COPY goAPI/go.sum ./
 
RUN go mod download

COPY goAPI/main.go goAPI/main.go

WORKDIR /app/goAPI
RUN go build -o goAPI
EXPOSE 8080

CMD [ "/app/goAPI/goAPI" ]