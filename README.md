# Tracing Responsability with Blockchain - Server (Under Development)

Tracing responsibility in _UML_ Diagrams using *Hiperledger Fabric* (_minifab_).

## Installation

Download the code and check the *fabric_commands.md*

## Commands

### Enter Vagrant VM
vagrant up
vagrant ssh

### Moves to Vagrant synced directory
cd /vagrant/

### Initialize the project
./commands/init.sh

### Push changes to GitHub and Blockchain Network
./commands/push.sh

### Pull a commit giving a commit SHA from GitHub
./commands/pull.sh

### Resets commit navigation to head
./commands/head.sh

### Clears the project
./commands/clear.sh

### Get history of transactions (Prints commits and author to console)
./commands/gethistory.sh

### Get history of transactions (Prints commits and author to console & outputs JSON files with block details)
./commands/gethistory.sh -out
