---
title: "Backup"
description: ""
lead: ""
date: 2022-10-31T16:33:44+08:00
lastmod: 2022-10-31T16:33:44+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "backup-3ea33cd3ba5204cb0f0a2c5d363c14fb"
weight: 999
toc: true
---
## Concepts
`Cold Backup` Full backup when database server is not running, less frequent. 

`Hot Backup` Full backup when database server is running and serving.

`Warn Backup` Full backup when database server is running, but locking some tables.

`Incremental Backup` Partial backup when database server is running, containing only data(updates) since the last backup. 

`Binlog` Binary format logs containing data updates. **Non-Idempotent**(different from MongoDB).

## Slave Backup
Slave server can be used as an efficient and safe way of replicating (not backup) data, but with minimal latency. 
Also, we can back up data in slave server, instead of master server.

A minimal master-slave example(mysql 8.0):

host
```shell
export dir=<>
mkdir -p ${dir}/master/mysql ${dir}/slave/mysql
docker run --name=mysql -p 3307:3306  -v ${dir}/master/mysql:/var/lib/mysql -e REPLICATION_MASTER=true  -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8.0
docker run --name=mysql-slave  --link mysql:mysql -p 3308:3306 -v ${dir}/slave/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456  -d mysql:8.0
```
master
```mysql
CREATE USER 'repl'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';
flush privileges;
show master status;  -- show current binlog file and position
```
slave
```mysql
set global server_id=2 -- non-persistent setting only for showcase
CHANGE MASTER TO 
  MASTER_HOST='172.17.0.3',
  MASTER_USER='repl',
  MASTER_PASSWORD='123456',
  MASTER_PORT=3306,
  MASTER_LOG_FILE='binlog.000002',            
  MASTER_LOG_POS=1324;
show slave status;
```
## Mysqldump & Binlog
mysqldump will dump data with binlog info, which can be helpful when restoring data with combined dump&binlog.
```shell
mysqldump --all-databases --source-data=2 --single-transaction > master-dump.sql
```
```mysql
CHANGE MASTER TO MASTER_LOG_FILE='binlog.000021', MASTER_LOG_POS=157;
```
mysqlbinlog can read directly from remote server and behaves like a non-serving slave server.
```shell
mysqlbinlog --read-from-remote-server -u root -p --raw --stop-never binlog.000016
```
Use dump file and binlog to restore data.
```shell
mysql -u root -p < master-dump.sql
mysqlbinlog -D -u root -p --start-position <MASTER_LOG_POS> [list of binlogs since MASTER_LOG_FILE]
```
## Reference
[MySQL数据库-binlog日志备份与增量恢复](https://devopssec.gitee.io/blog/2018/08/28/MySQL%E6%95%B0%E6%8D%AE%E5%BA%93-binlog%E6%97%A5%E5%BF%97%E5%A4%87%E4%BB%BD%E4%B8%8E%E5%A2%9E%E9%87%8F%E6%81%A2%E5%A4%8D/)

[优雅地给正在运行的 MySQL 添加从库](https://zahui.fan/posts/86a9c8f5/)

[MySql 主库/从库](https://www.jianshu.com/p/6cd5d57c6d1d)

[mysqlbinlog-backup](https://dev.mysql.com/doc/refman/8.0/en/mysqlbinlog-backup.html)

[recover-deleted-database](https://developer.aliyun.com/article/495602)
