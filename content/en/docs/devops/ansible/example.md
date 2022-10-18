---
title: "Example"
description: ""
lead: ""
date: 2022-09-12T18:27:34+08:00
lastmod: 2022-09-12T18:27:34+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "example-a36368ea51576a34e33e027b91e8d8de"
weight: 999
toc: true
---
```yaml
---
- name: "k8s ops"
  # Create a hosts file , and use -i <host_file>
  hosts: k8s
  become: yes
  become_user: yi.wu
  remote_user: yi.wu
  tasks:
    - name: "Playbook"
      # Run shell command
      shell: "sudo docker pull alpine:3.16"
      # Command args
      args:
        # Use bash
        executable: "/bin/bash"
      # Capture output
      register: output

    - debug: var=output.stdout_lines

```