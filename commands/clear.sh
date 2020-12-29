#!/bin/bash

# Changing directories
cd /var/lib/waagent/custom-script/download/0/project/bloc-server/

# Tearing down the blockchain network
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab down -o faculty.com

# Cleaning up the blockchain network
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab cleanup -o faculty.com



