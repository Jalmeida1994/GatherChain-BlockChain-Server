# Commands to Know!

### Moves to Vagrant synced directory
cd /vagrant/

### Up the Network
/vagrant/minifab up -o student.com

### Down the Network
/vagrant/minifab down

### Network up after its down
/vagrant/minifab netup

### Cleanup minifab
/vagrant/minifab cleanup

## Create Channel
/vagrant/minifab create -c testchannel

/vagrant/minifab join

/vagrant/minifab channelquery

### You can now make changes to config.json

/vagrant/minifab channelsign

/vagrant/minifab channelupdate

/vagrant/minifab channelquery

## Chaincode
/vagrant/minifab install -n testcc -l go

/vagrant/minifab approve

/vagrant/minifab commit -p "init","42490-43528","7bx89","42490"

### Instantiation of the chaincode
/vagrant/minifab instantiate -n testcc -p '"init","42490-43528","7bx89","42490"'

### Invoking (Committing the git commit)
/vagrant$ /vagrant/minifab invoke -n testcc -p '"invoke","42490-43528","89fd04","43528"'

### Query (Query the chaincode)
/vagrant/minifab invoke -n testcc -p '"query","42490-43528"'

## Block Query
/vagrant/minifab blockquery

## Helper Commands
### Check the files permissions
ls -l

### Give all permissions to file x.sh
chmod 755 x.sh