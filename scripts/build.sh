#!/bin/bash

set -e

release_path="builds"
targets=${@-"darwin/amd64 darwin/386 linux/amd64 linux/386 windows/amd64 windows/386"}
mkdir -p $release_path

for target in $targets; do
    os="$(echo $target | cut -d '/' -f1)"
    arch="$(echo $target | cut -d '/' -f2)"
    output="${release_path}/sup_${os}_${arch}"
    GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o $output
    zip -j $output.zip $output > /dev/null
done

cd $release_path
ls -al
