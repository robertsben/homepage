#!/bin/bash

set -e

for expected in GCLOUD_SERVICE_KEY GOOGLE_PROJECT_ID GOOGLE_COMPUTE_ZONE; do
    if [ -z "${!expected+x}" ]; then
        echo "env var $expected is not defined"
        exit 1
    fi
done


echo $GCLOUD_SERVICE_KEY | gcloud auth activate-service-account --key-file=-
gcloud --quiet config set project ${GOOGLE_PROJECT_ID}
gcloud --quiet config set compute/zone ${GOOGLE_COMPUTE_ZONE}
gcloud auth configure-docker
