---
title: "Bash"
description: ""
lead: ""
date: 2022-09-11T21:06:21+08:00
lastmod: 2022-09-11T21:06:21+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "bash-8d5786dbeb4cca7999494049a9c46611"
weight: 999
toc: true
---

script
```bash
## sum with awk
awk '{sum += $2} END {print sum}'

## ls sort by file size
ls -lh . --sort=size

## git 
git config --global https.proxy http://127.0.0.1:1086
git config --global https.proxy https://127.0.0.1:1086
git config --global http.proxy socks5://127.0.0.1:1086
git config --global https.proxy socks5://127.0.0.1:1086
git config --global --unset http.proxy
git config --global --unset https.proxy

## count lines
wc -l 'find . -name "*.java*"'
```

system 
```bash
## IO
iostat -xdm 1    

## CPU
mpstat -P ALL
lscpu | grep cpu
cat /proc/cpuinfo | grep processor

## memory
pmap <pid>
cat /proc/<pid>/status

## os
cat /proc/version
uname -a
lsb_release -a

## user
useradd wuyi
passwd wuyi
usermod -aG sudo wuyi
# add sudoers
sudo visudo
```