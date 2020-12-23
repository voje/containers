FROM alpine

RUN apk update && apk add py3-pip

RUN apk add python3-dev
RUN apk add libffi-dev
RUN apk add libressl-dev
RUN apk add build-base

RUN pip3 install openstackclient

RUN apk add vim curl nmap

WORKDIR /openstackrc

CMD ["/bin/sh"]
