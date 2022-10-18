---
title: "Localhost"
description: ""
lead: ""
date: 2022-09-12T18:26:04+08:00
lastmod: 2022-09-12T18:26:04+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "localhost-6a74763f294a02f1948b190e1dea9a32"
weight: 999
toc: true
---
```yaml
---
- name: "Playing with Ansible"
  hosts: localhost
  connection: local
  tasks:

  - name: "ls -l"
    shell: "hostname"
    register: "output"

  - debug: var=output.stdout_lines
```