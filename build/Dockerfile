FROM golang:1.17

WORKDIR /app

ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ADD . .

RUN go build -mod vendor -o scraper cmd/main.go

EXPOSE 3000

CMD [ "./scraper" ]