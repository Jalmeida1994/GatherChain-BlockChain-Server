#!/bin/bash

printf "Author: ${1}, group: ${2}, commit: ${3}"

cd /vagrant/

# Creates a channel in the network with the args given
/vagrant/minifab create -c channel${2} -o student.com 

# Joins the peers to the channel
/vagrant/minifab join

# Install, approve, commit and initialize our own chaincode
/vagrant/minifab approve,commit,initialize,discover -n testcc -l go -p \"init\",\"${2}\",\"${3}\",\"${1}\"


