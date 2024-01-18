#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

USER_ID=$(id -u)
GROUP_ID=$(id -g)

DOCKER_OPTS=${DOCKER_OPTS:-""}

# Build all iam commands.
function iam::build::build_command() {
  iam::log::status "Running build command..."
  make -C "{IAM_ROOT}" build.multiarch BINS="iamctl iam-apiserver iam-authz-server iam-pump iam-watcher"
}