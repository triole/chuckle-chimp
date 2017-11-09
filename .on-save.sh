#!/bin/bash

name="chuckjoke"

# architectures to build
archs=(
   "amd64"
   "arm"
)

# vars
builddir="./aedificatio/"
mkdir -p ${builddir}

# basic datections
# scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/"
# maingo=($(find ${scriptdir} -maxdepth 1 -type f -regex ".*\.go$"))
# name=$(echo ${maingo} | grep -Po "([^/]+$)" | grep -Po ".*?(?=\.)")

go-bindata data/
go fmt ${maingo}

for a in "${archs[@]}"; do
   tf="${builddir}${name}_${a}"
   GOOS=linux GOARCH=${a} go build -ldflags="-s -w" -o ${tf} *.go
   if [[ $? != "0" ]]; then
      notify-send "Go building \"${tf}\" failed."
      break
   fi
done

# AUTO ARCH DETECTION
# arch=$(dpkg --print-architecture)
# if [[ $? == "0" ]]; then
#    go build -ldflags="-s -w" -o ${builddir}${name}_${arch} ${gofile}
# else
#    notify-send "Go build failed."
# fi
