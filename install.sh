#!/bin/sh

BINARY='/usr/local/bin'

echo "Building dateseq"
go build dateseq.go

echo "Installing dateseq to $BINARY"
install -v dateseq $BINARY

echo "Removing the build"
rm dateseq
