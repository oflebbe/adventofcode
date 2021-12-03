#!/bin/sh

gamma=$(for i in 1 2 3 4 5 6 7 8 9 10 11 12 ; do
  cut -b $i input.txt | sort | uniq -c | sort -nr | head -1 | cut -f3 -d" "
done | tr -d "\n" ) 
dec=$(echo "ibase=2; $gamma"|bc)
inv=$(echo "2^12-1-$dec"|bc)
echo "$dec*$inv" | bc

#--------


oxyd=""
cat input.txt >tmp
for i in 1 2 3 4 5 6 7 8 9 10 11 12 13; do
    grep "^$oxyd" tmp| sort > tmp.2
    if [ `wc -l < tmp.2` = 1 ] ; then
        break
    fi
    res=`cut -b$i-$i <tmp.2 | sort | uniq -c| sort -rn`
    if [ `echo "$res" | awk '{print $1}' | uniq | wc -l` = 1 ] ; then
        oxyd="$oxyd"1
    else
        b=`echo "$res" | awk '{print $2; exit 0}'`
        oxyd="$oxyd$b"
    fi
    mv tmp.2 tmp
done

echo "ibase=2; `cat tmp.2`"|bc > tmp
oxyd=`cat tmp`

co2=""
cat input.txt >tmp
for i in 1 2 3 4 5 6 7 8 9 10 11 12 13; do
    grep "^$co2" tmp| sort > tmp.2
    if [ `wc -l < tmp.2` = 1 ] ; then
        break
    fi
    res=`cut -b$i-$i <tmp.2 | sort | uniq -c| sort -n`
    if [ `echo "$res" | awk '{print $1}' | uniq | wc -l` = 1 ] ; then
        co2="$co2"0
    else
        b=`echo "$res" | awk '{print $2; exit 0}'`
        co2="$co2$b"
    fi
    mv tmp.2 tmp
done

echo "ibase=2; `cat tmp.2`"|bc > tmp
co2=`cat tmp`

echo "$co2*$oxyd" | bc