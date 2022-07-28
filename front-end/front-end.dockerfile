FROM alpine:latest

RUN mkdir /app

COPY frontEndApp /App

CMD ["/app/frontEndApp"]