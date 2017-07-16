FROM alpine:latest

MAINTAINER Pedro Brandão <pcrbrandao@gmail.com>

WORKDIR "/opt"

ADD .docker_build/pet-amigo-server /opt/bin/pet-amigo-server
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/pet-amigo-server"]

