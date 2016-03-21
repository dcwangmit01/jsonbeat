#!/bin/bash
# set -x

# settings
CONTAINER=${CONTAINER:-jsonbeat}
VERSION=${VERSION:-0.1.0}
REGISTRY=${REGISTRY:-"gcr.io"}
PUSH=${PUSH:-true}

# program
PROJECT=${PROJECT:-`gcloud config list 2> /dev/null | grep project | cut -d'=' -f2 | xargs`}
COMMIT=${COMMIT:-`git log -1|grep commit|awk '{print $2}'|xargs`}
SHORT_COMMIT=${SHORT_COMMIT:-`echo "${COMMIT}"|head -c 7`}
TAG=${TAG:-"${VERSION}-$SHORT_COMMIT"}
IMAGE=${IMAGE:-"${REGISTRY}/${PROJECT}/${CONTAINER}:${TAG}"}

# build
docker build -f Dockerfile -t ${IMAGE} .

# push
if [ ${PUSH} = true ] ; then
    gcloud docker push ${IMAGE}
fi

