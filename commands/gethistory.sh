#!/bin/bash

cd /vagrant/

#Gets history of transactions for the group
/vagrant/minifab invoke -n testcc -p \"history\",\"${1}\" -o faculty.com

if [ "$1" == "-out" ]; then

find /vagrant/vars/ -name 'testchannel_*_txs.json' -exec rm {} \;

# Query the network for the newest block
/vagrant/minifab blockquery -b newest

# Takes the number of the newest block
nr_blocks=$(jq -r '.number' /vagrant/vars/testchannel_newest_txs.json)

# Removes 1 to total so not to query again same block
((nr_blocks--))

# Queries all blocks from the network, from the newest to the first (0)
until [ $nr_blocks -eq 5 ]
    do
        /vagrant/minifab blockquery -b $nr_blocks
        ((nr_blocks--))
    done

echo "Block details JSON files output format: /vagrant/vars/testchannel_*_txs.json"

fi