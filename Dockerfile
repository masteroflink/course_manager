FROM golang:1.20.2

# RUN apk update
# RUN apk add --no-cache git

WORKDIR /code

COPY go.* .
RUN go mod download

COPY . .
COPY .env .env

EXPOSE 8080

RUN go build -o /course-manager

CMD ["/course-manager"]
