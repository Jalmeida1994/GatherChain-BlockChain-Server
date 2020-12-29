# Commands to Know!

### Moves to working directory
cd /var/lib/waagent/custom-script/download/0/project/bloc-server/

### Up the Network
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab up -o student.com

### Down the Network
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab down

### Network up after its down
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab netup

### Cleanup minifab
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab cleanup

## Create Channel
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab create -c testchannel

/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab join

/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab channelquery

### You can now make changes to config.json

/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab channelsign

/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab channelupdate

/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab channelquery

## Chaincode
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab install -n testcc -l go

/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab approve

/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab commit -p "init","42490-43528","7bx89","42490"

### Instantiation of the chaincode
/var/lib/waagent/custom-script/download/0/project/bloc-server/minifab instantiate -n testcc -p '"init","42490-43528","7bx89","42490"'

### Invoking (Committing the git commit)
/var/lib/waagent/custom-script/download/0/project/bloc-server//minifab invoke -n testcc -p '"invoke","42490-43528","89fd04","43528"'

### Query (Query the chaincode)
/var/lib/waagent/custom-script/download/0/project/bloc-server//minifab invoke -n testcc -p '"query","42490-43528"'

## Block Query
/var/lib/waagent/custom-script/download/0/project/bloc-server//minifab blockquery

## Helper Commands
### Check the files permissions
ls -l

### Give all permissions to file x.sh
chmod 755 x.sh