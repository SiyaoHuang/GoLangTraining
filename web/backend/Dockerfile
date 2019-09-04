FROM golang:1.12.9-alpine3.10
RUN apk add bash ca-certificates git gcc g++ libc-dev
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN ls
RUN go build -o backend
RUN ls
CMD ["/app/backend"]
