#!/bin/bash

cd /var/lib/waagent/custom-script/download/0/project/bloc-server/

#Gets history of transactions for the group
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab invoke -n testcc -p \"history\",\"${1}\" -o faculty.com

if [ "$1" == "-out" ]; then

find /var/lib/waagent/custom-script/download/0/project/bloc-server/vars/ -name 'testchannel_*_txs.json' -exec rm {} \;

# Query the network for the newest block
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab blockquery -b newest

# Takes the number of the newest block
nr_blocks=$(jq -r '.number' /var/lib/waagent/custom-script/download/0/project/bloc-server/vars/testchannel_newest_txs.json)

# Removes 1 to total so not to query again same block
((nr_blocks--))

# Queries all blocks from the network, from the newest to the first (0)
until [ $nr_blocks -eq 5 ]
    do
        /var/lib/waagent/custom-script/download/0/project/bloc-server/minifab blockquery -b $nr_blocks
        ((nr_blocks--))
    done

echo "Block details JSON files output format: /var/lib/waagent/custom-script/download/0/project/bloc-server/vars/testchannel_*_txs.json"

fi