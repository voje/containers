FROM python:3.7-alpine as builder

RUN apk --no-cache add \
    build-base \
    python3-dev \
    libffi-dev \
    openssl-dev \
    ncurses-dev \
    libxml2-dev \
    libxslt-dev \
    git \
    bash \
    cargo \
    rust && \
    # netconf-console installation
    pip3 install git+https://bitbucket.org/martin_volf/ncc/@2.3.0 && \
    find /usr/lib/ -name '__pycache__' -print0 | xargs -0 -n1 rm -rf && \
    find /usr/lib/ -name '*.pyc' -print0 | xargs -0 -n1 rm -rf


FROM python:3.7-alpine as prod

LABEL LABEL project=https://github.com/voje/containers

RUN apk add --no-cache \
    python3 \
    openssh && \
    # cleanup
    find /usr/lib/ -name '__pycache__' -print0 | xargs -0 -n1 rm -rf && \
    find /usr/lib/ -name '*.pyc' -print0 | xargs -0 -n1 rm -rf

# Quality of life packages
RUN apk add --no-cache \
    vim \
    tmux

COPY --from=builder /usr/lib/libx*.* /usr/lib/
COPY --from=builder /usr/lib/libex*.* /usr/lib/
COPY --from=builder /usr/lib/libgcrypt* /usr/lib/
COPY --from=builder /usr/lib/libgpg* /usr/lib/
COPY --from=builder /usr/local/lib/python3.7/site-packages/ /usr/local/lib/python3.7/site-packages/
COPY --from=builder /usr/local/bin/netconf-console /usr/local/bin/netconf-console

WORKDIR /root

CMD [ "netconf-console" ]
