#!/bin/sh

echo "Dumping environment variables..."
echo
env
echo "---------------------------------"
echo "Content of GITHUB_OUTPUT"
cat $GITHUB_OUTPUT
echo "---------------------------------"
echo "Running scan..."
/scan
echo "---------------------------------"
