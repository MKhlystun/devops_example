FROM golang:latest
RUN addgroup --gid 10001 app
RUN adduser --gid 10001 --uid 10001 \
    --home /app --shell /sbin/nologin \
    --disabled-password --gecos app app

COPY bin/main /app/main
USER app
EXPOSE 8080
WORKDIR /app
ENTRYPOINT /app/main
