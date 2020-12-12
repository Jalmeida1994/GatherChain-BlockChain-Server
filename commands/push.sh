#!/bin/bash

cd /vagrant/

printf "Author: ${1}, group: ${2}, commit: ${3}"

#Pushes to the blockchain network
/vagrant/minifab invoke -c channel${2} -p \"invoke\",\"${2}\",\"${3}\",\"${1}\" -o student.com
