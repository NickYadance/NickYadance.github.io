---
layout: post
title:  "Mysql基本设置"
categories: Mysql
date:   2018-11-7 17:20:54
tags:  Mysql
---

* content
{:toc}

![mysqllogo.jpg](https://nickyadance.github.io/img/mysql.png)

## utf8 编码设置
停止mysql服务，打开文件

    /etc/my.cnf

添加如下字段：

```
[mysqld]
collation-server = utf8_unicode_ci
init-connect='SET NAMES utf8'
character-set-server = utf8

[client]
default-character-set=utf8

[mysql]
default-character-set=utf8
```

## 远程非root用户连接
创建用户：玩具程序host通常设置为'%'，允许所有ip登录
    
    CREATE USER 'username'@'host' IDENTIFIED BY 'password';
    
分配权限：个人习惯于分配单个数据库权限

    GRANT ALL ON DB.* to 'username'@'host' identified by 'password';
    
