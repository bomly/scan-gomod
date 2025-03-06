#!/usr/bin/env bash

echo "Dumping environment variables..."
echo
env
echo "Content of GITHUB_ENV"
cat $GITHUB_ENV
echo
