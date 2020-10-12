#!/usr/bin/env bash

set -x
# check if thrift exist

if ! command -v thrift &> /dev/null
then
    echo "thrift could not be found"
    exit
fi

# clone Evernote thrift definition
git clone --depth=1 https://github.com/evernote/evernote-thrift

# generate code
for file in ./evernote-thrift/src/*.thrift;
do
    thrift -gen go:package=evernote -o ./pkg ${file}
done    

# clean up
rm -rf evernote-thrift
