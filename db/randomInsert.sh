#!/bin/bash

insertNum=$1
maxStrLen=20
i=0
randomStr=zifuchuan
out=exe.sql
outTable=pm
nowRes=tmpStr
genRandomStr(){
		i=$RANDOM
		let i=i%maxStrLen	
		randomStr=$(openssl rand -base64 $i)
		nowRes=$randomStr
}
rm exe.sql

for nowInsertNum in $(seq 1 $insertNum)
	do
	
		genRandomStr	
		tmp1=$nowRes
		genRandomStr
		tmp2=$nowRes
		genRandomStr
		tmp3=$nowRes

		echo 'insert into' $outTable ' (b, c, d) values ("'$tmp1'",'$RANDOM',"'$tmp2'");'>>$out 

	done
