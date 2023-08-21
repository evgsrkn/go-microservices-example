FROM alpine:latest

WORKDIR /app
COPY . /app

RUN apk add gcompat

EXPOSE 8085

ENTRYPOINT [ "/app/bin/task" ]
