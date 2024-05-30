#!/bin/bash

FILES=dictionary-new/*

for f in $FILES
do
	#echo ${f##*/}
	VENDOR=${f#*.}

	mkdir $VENDOR
	./radius-dict-gen -package $VENDOR $f > $VENDOR/generated.go
done
