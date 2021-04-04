#!/bin/bash

if [ -e .envrc ]; then
	source .envrc
fi

pushd motting/api
go test -v
popd

pushd motting/cmd
go test
popd

pushd motting/dbaccess
go test -v
popd

pushd motting/model
go test
popd

pushd motting/server
go test
popd