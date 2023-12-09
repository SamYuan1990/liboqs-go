#!/bin/bash
get_arch=`uname -m`
if [[ $get_arch =~ "x86_64" ]];then
    wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
else
    wget https://go.dev/dl/go1.21.5.linux-$(uname -m).tar.gz
fi