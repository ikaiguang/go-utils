#!/usr/bin/env bash

# load your env
# source /etc/bashrc
# source /etc/profile
# source ~/.bashrc
# source ~/.profile
# source ~/.bash_profile

# ca.key
openssl genrsa -out ca.key 4096

# client.crt common_name=uufff.com
openssl req -new -x509 -days 36500 -key ca.key -out client.crt

# server.key
openssl genrsa -out server.key 4096

# server.csr common_name=uufff.com password=
openssl req -new -key server.key -out server.csr

# server.crt
openssl x509 -req -days 36500 -in server.csr -CA client.crt -CAkey ca.key -CAcreateserial -out server.crt
