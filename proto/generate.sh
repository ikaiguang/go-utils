#!/usr/bin/env bash

# load your env
# source /etc/bashrc
# source /etc/profile
# source ~/.bashrc
# source ~/.profile
# source ~/.bash_profile

protoc -I. -I$GOPATH/src --go_out=. ./*.proto
