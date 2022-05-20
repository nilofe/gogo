FROM golang:1.17-alpine

WORKDIR  /app
COPY go.mod ./


RUN go mod download 
RUN go mod verify
COPY *.go ./
RUN go build -o /gogo
EXPOSE 8080
CMD ["/gogo"]