#!/bin/bash

cd /var/lib/waagent/custom-script/download/0/project/bloc-server/

printf "Author: ${1}, group: ${2}, commit: ${3}"

#Pushes to the blockchain network
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab invoke -c channel${2} -p \"invoke\",\"${2}\",\"${3}\",\"${1}\" -o student.com
