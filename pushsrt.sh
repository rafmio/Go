#!/bin/bash
git pull
STATUS=$(git stataus)
UPTODATE="nothing to commit, working tree clean"
if [[ $STATUS == *"$UPTODATE"* ]]; then

DATETIME=$(date +'%d.%m.%Y %H:%M')

echo "printing $DATETIME"
