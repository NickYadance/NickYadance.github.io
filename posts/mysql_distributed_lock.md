---
title: 'Mysql分布式锁'
date: '2024-01-11'
description: '尽管分布式锁有很多成熟的实现方案，Mysql却常常是小应用场景下的第一存储选择'
---

最近在一份代码中看到基于Mysql的分布式锁实现，首先有一张锁表(脱敏)：
```mysql
CREATE TABLE `lock_tab`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `lock_type`  int                 NOT NULL,
    `lock_owner` varchar(127)        NOT NULL,
    `lock_ts`    datetime            NOT NULL,
    `reserved`   text default NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_idx_lock_type` (`lock_type`)
)
```

加锁流程的伪代码
```go
// distributed lock with Mysql
hostIp := 192.168.0.1
lockType := 1
lock := db.model("lock_tab").where("lock_type = ?", some_lock_type)
if !lock {
	db.model("lock_tab").insert(&Lock{LockType: lockType, LockOwner: hostIP, LockTs: time.Now()})
	return true
}

// the lock belongs to someone else and not expired yet
if lock.LockOwner != hostIp && notExpire(lock.Ts) {
	return false
}

// lock acquired, refresh the lease
db.model("lock_tab").update(&Lock{LockType: lockType, LockOwner: hostIP, LockTs: time.Now()})
return true
```

这份分布式锁的实现有几个特点：
1. 以进程所在机器(容器)的ip作为标识，因此每台机器最多一个实例获取到锁
2. 带过期时间
3. 无锁，没有施加行锁

这份代码的本意是限制crontab任务在多实例上的运行，因此才有lockType对应不同的cronTab任务。那么这份代码有没有问题，能不能正常工作呢？

我的答案是**有很大问题，但大概率能正常工作**。问题在于当锁过期时，有很高的风险会出现多个实例获取到锁并错误的对锁做续期，又由于实例都是以crontab周期获取锁，当实例数超过一定值问题出现的频率会非常高。但这份代码的运行场景恰好只有2个实例，因此大概率是能正常工作。同时，这份代码还存在一些时间戳同步、ip重用的细节问题。

这份代码也引起我对一个常见问题的思考：到底如何基于Mysql实现一份可用、高效的分布式锁呢？

在Redis中，实现一个分布式锁是简单而直观的：
```
// 加锁
SET my_lock my_random_value NX PX 30000

// 解锁
EVAL "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end" 1 my_lock my_random_value
```

以Redis为标杆，我对Mysql分布式锁的实现思路有：
1. server side，规避本地计算可能出现的时间戳问题
2. 原子性
3. 无锁

在一番搜索和对比后chatgpt4给出的解是最符合要求的，同时兼容Mysql5.0/8.0。实际运用中，程序可以生成一个随机值作为`lock_name`，判断加锁语句返回的affectedRows来判定是否加锁成功。
```mysql
CREATE TABLE `distributed_lock` (
  `lock_name` VARCHAR(255) NOT NULL,
  `lock_timestamp` BIGINT UNSIGNED NOT NULL,
  `lock_ttl` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`lock_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# 加锁
INSERT INTO `distributed_lock` (`lock_name`, `lock_timestamp`, `lock_ttl`)
VALUES ('my_lock', UNIX_TIMESTAMP(), 10)
ON DUPLICATE KEY UPDATE
                     `lock_timestamp` = IF(UNIX_TIMESTAMP() - `lock_timestamp` > `lock_ttl`, VALUES(`lock_timestamp`), `lock_timestamp`),
                     `lock_ttl` = IF(UNIX_TIMESTAMP() - `lock_timestamp` > `lock_ttl`, VALUES(`lock_ttl`), `lock_ttl`);

# 解锁
DELETE FROM `distributed_lock` WHERE `lock_name` = 'my_lock';
```
