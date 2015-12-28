#!/bin/sh
set -eu
echo "> VERSION"
go version
echo "> TEST"
go test
echo "> BUILD"
go build
echo "> SUCCEEDED"
