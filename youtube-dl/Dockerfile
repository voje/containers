FROM python:3.12.0-alpine3.18

ARG short_commit

RUN apk add git make zip pandoc

WORKDIR /opt
RUN git clone https://github.com/ytdl-org/youtube-dl.git && \
    cd youtube-dl && \
    git checkout -b $short_commit && \
    git log -n 3 && \
    rm -r .git

WORKDIR /opt/youtube-dl
RUN make install && youtube-dl --version

WORKDIR /project
