#!/bin/bash

read -p 'Student Number: ' username
read -p 'Password: ' password

. app.env

#Checks if user or pass are empty
if [ -z "$username" ] || [ -z "$password" ]
then
      echo "ERROR: Cannot login with empty values"
      exit 1  
else

if [ -f login.env ]; then
    . login.env

if [ -n "$LOGINUSER" ] || [ -n "$LOGINPASS" ]
then
    head -n -2 login.env > tmp.txt && cp tmp.txt login.env
    unset LOGINUSER
    unset LOGINPASS
fi
fi

echo "export LOGINUSER=${username}" >> login.env
echo "export LOGINPASS=${password}" >> login.env

fi

