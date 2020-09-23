FROM debian:stretch

RUN apt-get update && apt-get install -y \
	aptly \
	curl \
	wget \
	procps \
	vim \
	gnupg

# rsync (requires ssh client)
RUN apt-get install -y \
	rsync \
	openssh-client \
	ca-certificates \
	&& update-ca-certificates
