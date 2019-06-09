#!/bin/bash

set -e

for expected in GOOGLE_PROJECT_ID DOCKER_IMAGE_NAME DOCKER_IMAGE_TAG BUILD_DIRECTORY; do
    if [ -z "${!expected+x}" ]; then
        echo "env var $expected is not defined"
        exit 1
    fi
done

DOCKER_IMAGE_URL="gcr.io/${GOOGLE_PROJECT_ID}/${DOCKER_IMAGE_NAME}"

docker build -t ${DOCKER_IMAGE_URL}:${DOCKER_IMAGE_TAG} ${BUILD_DIRECTORY}
docker push ${DOCKER_IMAGE_URL}:${DOCKER_IMAGE_TAG}
