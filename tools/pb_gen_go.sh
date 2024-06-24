#!/bin/bash -eu
#
# Copyright 2017-2023 Google Inc.
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

# This script generates Go code for the config protobufs.

PROJECT="cloudprober1"
PROJECT_ROOT="${PWD}/../../"

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
    PROJECTROOT="${SCRIPTDIR}/../.."
  else
    echo "PROJECTROOT is not set and we are not able to determine PROJECTROOT"
    echo "from script's path. PROJECTROOT should be set such that project files "
    echo " are located at $PROJECT relative to the PROJECTROOT."
    exit 1
  fi
fi
echo PROJECTROOT=${PROJECTROOT}

function cleanup {
  echo "Removing temporary directory used for protoc installation: ${TMPDIR}"
  rm  -r "${TMPDIR}"
}
#trap cleanup EXIT

# Install cue from Go
go install cuelang.org/go/cmd/cue@latest

echo "Generating Go code and CUE schema for protobufs.."
echo "======================================================================"
# Generate protobuf code from the root directory to ensure proper import paths.
cd $PROJECTROOT

# Create a temporary director to generate protobuf Go files.
TMPDIR=$(mktemp -d)
echo $TMPDIR
mkdir -p ${TMPDIR}/github.com/cloudprober/cloudprober
rsync -mr --exclude='.git' --include='*/' --include='*.proto' --include='*.cue' --exclude='*' $PROJECT/* $TMPDIR/github.com/cloudprober/cloudprober

cd $TMPDIR

MODULE=github.com/cloudprober/cloudprober

echo "Generating CUE schema from protobufs.."
cd $MODULE
cue import proto -I . config/proto/config.proto --proto_enum json -f

# Generate Go code for proto
find ${MODULE} -type d | \
  while read -r dir
  do
    # Ignore directories with no proto files.
    ls ${dir}/*.proto > /dev/null 2>&1 || continue
    ${protoc_path} --go-grpc_out=. --go_out=. ${dir}/*.proto
  done

# Copy generated files back to their original location.
find ${MODULE} \( -name *.pb.go -o -name *proto_gen.cue \) | \
  while read -r pbgofile
  do
    dst=${PROJECTROOT}/${pbgofile/github.com\/cloudprober\//}
    cp "$pbgofile" "$dst"
  done

cd -
