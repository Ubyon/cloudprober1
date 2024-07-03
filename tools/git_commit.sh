#!/bin/bash

BRANCH=$(git symbolic-ref --short HEAD)

# Get the last commit ID (short form)
commit_id_short=$(git log -1 origin/$BRANCH --format=%H | cut -c1-12)

# Get the last commit timestamp in the desired format
commit_timestamp=$(git log -1 origin/$BRANCH --format=%cd --date=format:'%Y%m%d%H%M%S')

# Print both values
echo "Commit ID (short): $commit_id_short"
echo "Timestamp: $commit_timestamp"
echo "Go.mod relace format: $commit_timestamp-$commit_id_short"
