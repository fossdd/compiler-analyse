#!/bin/sh

echo du -h $@ > tmp.sh

IFS='	'
read value _ <<< $(bash tmp.sh 2>&1 | tail -n3 | head -n1)

printf $value

rm tmp.sh
