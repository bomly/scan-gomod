#!/bin/sh

echo "Dumping environment variables..."
echo
env
echo "---------------------------------"
echo "Running scan..."
result=$(/scan | base64)
echo "RESULT=${result}" >> $GITHUB_OUTPUT
echo "---------------------------------"
echo "Content of GITHUB_OUTPUT"
cat $GITHUB_OUTPUT
echo "---------------------------------"
echo "Displaying result..."
echo $result | base64 --decode
echo "---------------------------------"
