---
title: "Kubectl"
description: ""
lead: ""
date: 2022-09-08T18:00:43+08:00
lastmod: 2022-09-08T18:00:43+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "kubectl-05d05291d8c09012e29c7f1dbef15f83"
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

