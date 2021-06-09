#!/bin/bash

# Copy the chaincode files to the vars folder
cd /var/lib/waagent/custom-script/download/0/project/bloc-server/
cp -R chaincode vars

printf "Author: ${1}, group: ${2}, commit: ${3}"

chmod +x commands
chmod +x vars
cd /var/lib/waagent/custom-script/download/0/project/bloc-server/

# Starts the network with the args given
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab up -o faculty.com -c channel${2} -e 7300 -n testcc -l go -p \"init\",\"${2}\",\"${3}\",\"${1}\"




