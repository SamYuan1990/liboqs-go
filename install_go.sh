#!/bin/bash
get_arch=`uname -m`
if [[ $get_arch =~ "x86_64" ]];then
    wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
elif [[ $get_arch =~ "arrch64" ]];then
    wget https://go.dev/dl/go1.21.5.linux-arm64.tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-arm64.tar.gz
else
    wget https://go.dev/dl/go1.21.5.linux-$(uname -m).tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-$(uname -m).tar.gz
fi