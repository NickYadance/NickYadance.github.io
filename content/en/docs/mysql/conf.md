---
title: "Conf"
description: ""
lead: ""
date: 2022-11-22T10:58:40+08:00
lastmod: 2022-11-22T10:58:40+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "conf-b4fc02517be29832e9d5b4ba56658e96"
weight: 999
toc: true
---
Host
```text
Cpu: 64C
Mem: 256GB
Disk: SSD
```
Initialize conf
```text
[mysqld]
max_connections  = 1024
innodb_buffer_pool_size = 137438953472 # 128GB, 60-80% of the mem
innodb_buffer_pool_chunk_size = 1073741824 # 1G
innodb_buffer_pool_instances = 8 # 8-16 if innodb_buffer_pool_chunk_size >= 1G, else <8
innodb_flush_log_at_trx_commit = 0 # 0,1,2, take care the data safety
innodb_log_buffer_size = 1073741824 # 1G
innodb_log_file_size = 4294967296 # 4G
tmp_table_size = 134217728
max_heap_table_size = 134217728
slow_query_log = 1 
long_query_time = 5
binlog_expire_logs_seconds = 3600
innodb_page_cleaners = 4
innodb_read_io_threads = 16
innodb_write_io_threads = 16
```