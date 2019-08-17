FROM alpine
MAINTAINER younccat

WORKDIR /app

COPY .                          .

EXPOSE 3000

ENTRYPOINT ["./docker/scripts/entrypoint.sh"]
