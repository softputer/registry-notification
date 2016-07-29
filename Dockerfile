FROM ubuntu:trusty

MAINTAINER Jayson Ge <yjge@qingyuanos.com>

COPY ./registry-notification /

CMD ["/registry-notification"]
