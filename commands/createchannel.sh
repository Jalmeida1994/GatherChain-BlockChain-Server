#!/bin/bash

printf "Author: ${1}, group: ${2}, commit: ${3}"

cd /var/lib/waagent/custom-script/download/0/project/bloc-server/

# Creates a channel in the network with the args given
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab create -c channel${2} -o student.com 

# Joins the peers to the channel
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab join

# Install, approve, commit and initialize our own chaincode
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab approve,commit,initialize,discover -n testcc -l go -p \"init\",\"${2}\",\"${3}\",\"${1}\"


