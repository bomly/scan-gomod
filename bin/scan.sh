#!/bin/sh

echo "Dumping environment variables..."
echo
env
echo "---------------------------------"
echo "Content of GITHUB_ENV"
cat $GITHUB_ENV
echo "---------------------------------"
echo "Running scan..."
/scan
echo "---------------------------------"
