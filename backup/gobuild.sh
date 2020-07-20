#!/bin/bash
./buildandrun.sh

// goa design installation
env GO111MODULE=on go get -u goa.design/goa/v3/...@v3
// setting environment variables
export GOBIN=$PWD/bin
export PATH=$GOBIN:$PATH
export PATH=$PATH:/usr/local/go/bin

goa gen fruitshop/design
goa example fruitshop/design
// Go build to generate binary file
go build ./cmd/fruitshop
chmod +x ./fruitshop

file="data.db"

if [ -f $file ] ; then
    rm $file
fi



// executing the application
./fruitshop
