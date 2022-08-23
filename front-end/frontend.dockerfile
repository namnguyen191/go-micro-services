FROM alpine:latest

RUN mkdir /app

COPY frontAppLinux /app

CMD [ "/app/frontAppLinux" ]