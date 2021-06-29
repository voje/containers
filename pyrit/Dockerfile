FROM debian:stable

RUN apt-get update

RUN apt-get install -y git vim curl

RUN apt-get install -y python2
RUN ln -s /usr/bin/python2 /usr/bin/python

WORKDIR /opt/pip2
RUN curl https://bootstrap.pypa.io/2.7/get-pip.py --output get-pip.py
RUN python2 ./get-pip.py
RUN pip install \
    psycopg2-binary \
    scapy==2.3.2

RUN apt-get install -y \
	gcc \
	python2-dev \
	libssl-dev \
	zlib1g-dev \
	libpcap-dev

WORKDIR /root/git/
RUN git clone https://github.com/JPaulMora/Pyrit.git

WORKDIR /root/git/Pyrit
RUN git checkout f0f1913c645b445dd391fb047b812b5ba511782c
RUN ./setup.py clean
RUN ./setup.py build
RUN ./setup.py install

RUN pyrit -h

