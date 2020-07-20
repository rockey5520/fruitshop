#!/bin/bash

file="fruitshop.db"

if [ -f $file ] ; then
   rm $file
fi



// executing the application
$HOME/go/bin/swag init
go run main.go
