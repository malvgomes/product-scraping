FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD . .

RUN go build -o scraper cmd/main.go

EXPOSE 3000

CMD [ "./scraper" ]