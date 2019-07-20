FROM alpine:3.10

RUN mkdir /app
WORKDIR /app
ADD ./.out/gofoody-app /app/gofoody-app

CMD ["./gofoody-app"]
