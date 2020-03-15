#!/bin/bash

if [ -e .envrc ]; then
	source .envrc
fi

pushd motting
go build
popd

pushd webpush
go build
popd