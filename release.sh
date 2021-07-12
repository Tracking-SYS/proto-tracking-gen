#!/bin/bash
echo "\nCopy artifacts"
rm -rf go
cp -R ../proto-tracking/go ./
cp -R ../proto-tracking/docs ./
cp -R ../proto-tracking/desc ./
go mod tidy
