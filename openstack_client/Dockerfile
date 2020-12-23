FROM alpine

RUN apk update && apk add py3-pip

RUN apk add python3-dev
RUN apk add libffi-dev
RUN apk add libressl-dev
RUN apk add build-base
RUN apk add vim curl nmap

RUN pip3 install --upgrade pip
RUN pip3 install openstackclient
RUN pip3 install ansible

RUN ansible-galaxy collection install openstack.cloud

WORKDIR /openstackrc

CMD ["/bin/sh"]
