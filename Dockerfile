FROM golang:latest

RUN mkdir /app

WORKDIR /app

ADD . /app

COPY . . 

RUN go mod download

RUN CGO_ENABLED=0 

RUN go build -o main .

ENTRYPOINT [ "/app/main" ]