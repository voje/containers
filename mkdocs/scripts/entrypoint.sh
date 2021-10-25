#!/bin/bash

INPUT_DIR=/var/mkdocs/raw
SRV_DIR=/srv
MKDOCS_PROJECT_NAME=mkdocs
MKDOCS_PROJECT_DIR=${SRV_DIR}/${MKDOCS_PROJECT_NAME}
MKDOCS_DATA_DIR=${MKDOCS_PROJECT_DIR}/docs

# Create mkdoc project
cd ${SRV_DIR}
mkdocs new ${MKDOCS_PROJECT_NAME}

# Copy files into the project's datadir
cp ${INPUT_DIR}/* ${MKDOCS_DATA_DIR}/. -r

# Run mkdocs server
cd ${MKDOCS_PROJECT_DIR}
mkdocs serve -a 0.0.0.0:8000
