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

## Mysqldump & Binlog
**DON'T use mysqldump to busy databases** cuz the table locking can potentially hang all your requests for a long time.

mysqldump will dump data with binlog info, which is helpful when restoring data with combined dump&binlog.
```shell
mysqldump --all-databases --source-data=2 --single-transaction > master-dump.sql
## binlog info in master-dump.sql:
## CHANGE MASTER TO MASTER_LOG_FILE='binlog.000021', MASTER_LOG_POS=157;
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

## Percona Xtrabackup (Recommended)
Before using Percona, **make sure that the Percona version is compatible with mysql server**.

Backup
```shell
xtrabackup --backup \
--target-dir=/path/to/backup \
--socket=<socket path, use `show variables like '%socket%'` to find out> \
-H <host> \
-u <user> \
-p
```

Backup with compress and stream
```shell
xtrabackup --backup \
--target-dir=/var/lib/mysql/dump \
--stream=xbstream \
--compress \
--parallel=4 \
-H 127.0.0.1 \
-S /var/run/mysqld/mysqld.sock \
-u root \
-p | ssh user@remotehost "xbstream -x -C /path/to/backup"
```

Decompress before prepare
```shell
xtrabackup --decompress --target-dir=/path/to/backup
```

Prepare
```shell
xtrabackup --prepare --target-dir=/path/to/backup
```

Restore
```shell
xtrabackup --copy-back --target-dir=/path/to/backup --datadir=<data directory for new server>
```

## Slave Backup
A minimal master-slave example(mysql 8.0):

master, create repl user
```mysql
CREATE USER 'repl'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';
flush privileges;
show master status;  -- show current binlog file and position
```
slave, start database in docker
```shell
docker run --name mysql \
-v /mysql/conf.d:/etc/mysql/conf.d \
-v /mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=root \
-p 30360:3306 \
-d mysql:8.0
```

```mysql
set global server_id=2;

set global read_only=1;
set global super_read_only=1;

CHANGE MASTER TO
  MASTER_HOST='192.168.0.1',
  MASTER_USER='root',
  MASTER_PASSWORD='root',
  MASTER_PORT=30360,
  MASTER_LOG_FILE='binlog.029168',  ## Find binlog file from data directory binlog.index
  MASTER_LOG_POS=724; ## Find pos from binlog file size(in byte)

start slave;

show slave status;
## Expected result:
##       Slave_SQL_Running_State: Replica has read all relay log; waiting for more updates
```

## Reference
[MySQL数据库-binlog日志备份与增量恢复](https://devopssec.gitee.io/blog/2018/08/28/MySQL%E6%95%B0%E6%8D%AE%E5%BA%93-binlog%E6%97%A5%E5%BF%97%E5%A4%87%E4%BB%BD%E4%B8%8E%E5%A2%9E%E9%87%8F%E6%81%A2%E5%A4%8D/)

[优雅地给正在运行的 MySQL 添加从库](https://zahui.fan/posts/86a9c8f5/)

[MySql 主库/从库](https://www.jianshu.com/p/6cd5d57c6d1d)

[mysqlbinlog-backup](https://dev.mysql.com/doc/refman/8.0/en/mysqlbinlog-backup.html)

[recover-deleted-database](https://developer.aliyun.com/article/495602)
