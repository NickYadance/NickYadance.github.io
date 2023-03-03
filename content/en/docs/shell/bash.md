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

## cut first field
command | cut -f 1 -d '='
## cut seconds until last fields
command | cut -f 2- -d '='

## git
git config --global https.proxy http://127.0.0.1:1086
git config --global https.proxy https://127.0.0.1:1086
git config --global http.proxy socks5://127.0.0.1:1086
git config --global https.proxy socks5://127.0.0.1:1086
git config --global --unset http.proxy
git config --global --unset https.proxy

## count lines
wc -l 'find . -name "*.java*"'

## find largest files
find -type f -exec du -Sh {} + | sort -rh | head -n 20

## jq
TLS_CRT=$($kubectl_host get secret -n karmada-system karmada-webhook-cert --template='{{ index .data "tls.crt"}}' | tr -d '\n')
TLS_KEY=$($kubectl_host get secret -n karmada-system karmada-webhook-cert --template='{{ index .data "tls.key"}}' | tr -d '\n')
$kubectl_host get secret -n kruise-system kruise-webhook-certs -o json \
        | jq --arg TLS_KEY "$TLS_KEY" '.data["tls.key"] |= $TLS_KEY' \
        | jq --arg TLS_CRT "$TLS_CRT" '.data["tls.crt"] |= $TLS_CRT' \
        | jq 'del(.metadata.annotations)' \
        | jq 'del(.metadata.resourceVersion)' \
        | jq 'del(.metadata.uid)' \
        | kubectl apply -f -

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

mysql
```mysql
CREATE USER 'repl'@'%' IDENTIFIED BY '123456';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';
flush privileges;

show processlist ;
select * from INFORMATION_SCHEMA.PROCESSLIST where id = ''\G
kill mysql_thread_id;
show engine innodb status;
select trx_id, trx_state, trx_mysql_thread_id, trx_query from information_schema.innodb_trx;

show variables like "max_connections";
set global max_connections = 200;

show global status like "Com_select";  do sleep(10); show global status like "Com_select";

```

kubectl
```shell
ls -l /sys/fs/cgroup/memory | grep system.slice

# get po sort by age
kubectl get pod --sort-by=.metadata.creationTimestamp

# get po group by hostname
kubectl get po -o wide | awk '{print $7}' | sort | uniq -c

# port forward
kubectl port-forward svc/airflow-webserver 8080:8080 --namespace airflow

# upsert configmap
kubectl create configmap foo --from-file foo.properties -o yaml --dry-run | kubectl apply -f -
```
