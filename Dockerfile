FROM ubuntu:latest

WORKDIR /app

COPY chat-server ./
COPY .env ./

EXPOSE 80

ENTRYPOINT ["./chat-server"]
