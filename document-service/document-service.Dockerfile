FROM alpine:latest

RUN mkdir /app

COPY documentApp /app

CMD ["/app/documentApp"]