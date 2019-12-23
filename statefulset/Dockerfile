FROM ubuntu:16.04

RUN mkdir -p /app

WORKDIR /app

COPY ./statefulset-demo /app
COPY ./config.json /app

EXPOSE 8080

# This command runs your application, comment out this line to compile only
CMD ["./statefulset-demo"]

LABEL Name=statefulset-demo Version=0.0.1