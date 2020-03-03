#!/bin/bash

if [ -e .envrc ]; then
	source .envrc
fi

pushd motting/cmd
go test
popd

pushd motting/dbaccess
go test
popd

pushd motting/server
go test
popd