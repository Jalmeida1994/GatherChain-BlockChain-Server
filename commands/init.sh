#!/bin/bash

# Copy the chaincode files to the vars folder
cd /vagrant/
cp -R chaincode vars

printf "Author: ${1}, group: ${2}, commit: ${3}"

chmod +x commands
chmod +x vars
cd /vagrant/

# Starts the network with the args given
/vagrant/minifab up -o faculty.com -c channel${2} -n testcc -l go -p \"init\",\"${2}\",\"${3}\",\"${1}\"




