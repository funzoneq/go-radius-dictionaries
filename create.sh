#!/bin/bash

FILES=dictionary/*

for f in $FILES
do
	#echo ${f##*/}
	VENDOR=${f#*.}

	mkdir $VENDOR
	$GOPATH/bin/radius-dict-gen -package $VENDOR $f > $VENDOR/generated.go
done
