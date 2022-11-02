---
title: "Client"
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
create user
```mysql
CREATE USER 'repl'@'%' IDENTIFIED BY '123456';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';
flush privileges;
```
find out the latest deadlocks
```mysql
show engine innodb status;
```
show transactions
```mysql
select trx_id, trx_state, trx_mysql_thread_id, trx_query from information_schema.innodb_trx;
```
show processeslist
```mysql
SHOW FULL PROCESSLIST;
```
kill processes
```mysql
kill <mysql_thread_id>;
```
get&set variables
```mysql
show variables like "max_connections";
set global max_connections = 200;
```
count select qps
```mysql
show global status like "Com_select";  do sleep(10); show global status like "Com_select";
```
