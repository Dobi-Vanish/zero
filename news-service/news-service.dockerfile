FROM alpine:latest

RUN mkdir /app


COPY newsApp /app
COPY cmd/api/migrations /app/migrations
COPY example.env /app/example.env


WORKDIR /app

CMD [ "/app/newsApp"]