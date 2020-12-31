FROM golang:latest as builder

RUN mkdir /app

WORKDIR /app

COPY go.mod .

COPY go.sum . 

RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main 

FROM alpine 


COPY --from=builder /app/main /main

EXPOSE 8080 

ENTRYPOINT [ "/app/main" ]