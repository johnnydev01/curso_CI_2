FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go

RUN chmod +x main

COPY ./templates/ templates/

EXPOSE 8000


CMD [ "./main" ]
