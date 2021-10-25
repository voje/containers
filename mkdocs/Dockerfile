FROM ubuntu:20.04

RUN apt-get update && apt-get install -y \
	python3-pip

RUN pip install mkdocs

COPY scripts/ /usr/local/bin/scripts/
ENV PATH=${PATH}:/usr/local/bin/scripts

CMD ["entrypoint.sh"]
