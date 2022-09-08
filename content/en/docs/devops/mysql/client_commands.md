---
title: "Client Commands"
description: ""
lead: ""
date: 2022-09-08T17:39:35+08:00
lastmod: 2022-09-08T17:39:35+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "client_commands-34a6706721f1b4023ca70185de9af3e9"
weight: 999
toc: true
---
find out the latest deadlocks
```mysql
show engine innodb status;
```
show transactions
```mysql
select trx_id, trx_state, trx_mysql_thread_id, trx_query from information_schema.innodb_trx;
```
show processes
```mysql
show processlist;
```
kill processes
```mysql
kill <mysql_thread_id>;
```
