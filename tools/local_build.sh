#!/bin/bash -eu
#
# Copyright 2017 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script builds cloudprober locally. It expects protobuf's Go code to
# be already available (can be done using tools/gen_pb_go.sh).

PROTOC_VERSION="25.2"
PROJECT="cloudprober1"

GOPATH=$(go env GOPATH)

if [ -z "$GOPATH" ]; then
  echo "Go environment is not setup correctly. Please look at"
  echo "https://golang.org/doc/code.html to set up Go environment."
  exit 1
fi
echo GOPATH=${GOPATH}

if [ -z ${PROJECTROOT+x} ]; then
  # If PROJECTROOT is not set, try to determine it from script's location
  SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
  if [[ $SCRIPTDIR == *"$PROJECT/tools"* ]]; then
    PROJECTROOT="${SCRIPTDIR}/.."
  else
    echo "PROJECTROOT is not set and we are not able to determine PROJECTROOT"
    echo "from script's path. PROJECTROOT should be set such that project files "
    echo " are located at $PROJECT relative to the PROJECTROOT."
    exit 1
  fi
fi
echo PROJECTROOT=${PROJECTROOT}

if [ ! -d "${PROJECTROOT}" ];then
  echo "${PROJECT} not found under Go workspace: ${GOPATH}/src. Please download"
  echo " cloudprober source code from github.com/cloudprober/cloudprober and set it "
  echo "such that it's available at ${PROJECTROOT}."
  exit 1
fi

cd ${PROJECTROOT}
# Get all dependencies
echo "Getting all the dependencies.."
echo "======================================================================"
go get -t ./...

# Build everything
echo "Build everything..."
echo "======================================================================"
go build ./...

# Run tests
echo "Running tests..."
echo "======================================================================"
go test ./...

# Install cloudprober
echo "Build static cloudprober binary.."
echo "======================================================================"
CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ./cmd/cloudprober.go
