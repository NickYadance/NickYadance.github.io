---
title: "Kubectl"
description: ""
lead: ""
date: 2022-09-11T21:05:41+08:00
lastmod: 2022-09-11T21:05:41+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "kubectl-018fdf5fe90e72c19dd1bc9282305051"
weight: 999
toc: true
---
check cgroup memory
```shell
ls -l /sys/fs/cgroup/memory | grep system.slice
```
get po sort by age
```shell
kubectl get pod --sort-by=.metadata.creationTimestamp
```
get po group by hostname
```shell
kubectl get po -o wide | awk '{print $7}' | sort | uniq -c
```
port forward
```shell
kubectl port-forward svc/airflow-webserver 8080:8080 --namespace airflow
```


