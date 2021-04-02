
#!/bin/sh

echo time $@ > tmp.sh

IFS='	'
read _ value <<< $(bash tmp.sh 2>&1 | tail -n3 | head -n1)

printf $value

rm tmp.sh
