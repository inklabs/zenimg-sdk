#!/bin/bash

check_status() {
	if [ $? -ne 0 ]; then
		echo
		echo "$RES"
	else
		printf "\t[\e[01;32mOK\e[00m]"
	fi
}

echo -n "PHP:"
RES=$(php php/test.php)
check_status


echo
echo -n "Ruby:"
RES=$(ruby ruby/test.rb)
check_status

echo
echo -n "Python:"
RES=$(python python/test.py)
check_status

echo
echo -n "JS:"
RES=$(cd js && node test.js)
check_status

echo
