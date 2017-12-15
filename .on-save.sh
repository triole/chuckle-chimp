#!/bin/bash

# settings and architectures to build
builddir="./aedificatio/"
archs=("amd64" "arm")

# self setting vars
scriptdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/"
name=$(echo ${scriptdir} | grep -Po "/([^/]+)/?$" | tr -d "/")

mkdir -p ${builddir}

go-bindata data/
go fmt *.go

for a in "${archs[@]}"; do
   tf="${builddir}${name}_${a}"
   GOOS=linux GOARCH=${a} go build -ldflags="-s -w" -o ${tf} *.go
   if [[ $? != "0" ]]; then
      notify-send "Go building \"${tf}\" failed."
      break
   fi
done
