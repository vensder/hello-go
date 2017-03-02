FROM alpine:3.5

MAINTAINER vensder <vensder@gmail.com>

COPY webserver /usr/local/bin/
EXPOSE 8080
CMD ["webserver"]

