FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/pet-amigo-server /opt/bin/pet-amigo-server
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/go-getting-started"]

