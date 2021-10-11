#!/usr/bin/env bash

set -o errexit
set -o pipefail

pulumi config set aws:profile "${AWS_PROFILE}"
pulumi config set aws:region "${AWS_REGION}"
