#!/bin/bash
echo "\nCopy artifacts"
rm -rf go
cp -R ../proto-tracking/go ./
go mod tidy
