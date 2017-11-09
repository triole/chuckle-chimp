#!/bin/bash

builddir="./aedificatio/"

# architectures to build
archs=(
   "amd64"
   "arm"
)

# vars
mkdir -p ${builddir}
name=$(echo ${scriptdir} | grep -Po "/([^/]+)/?$" | tr -d "/")

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
