---
title: "Lock"
description: ""
lead: ""
date: 2022-10-11T18:36:36+08:00
lastmod: 2022-10-11T18:36:36+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "lock-edc12644128700d01c82c0073965e90c"
weight: 999
toc: true
---
{{< details "Setup" >}}
```mysql
CREATE DATABASE study;
USE study;

CREATE TABLE student_scores (
  id int(11) NOT NULL,
  name varchar(100) NOT NULL,
  age tinyint(4) NOT NULL,
  score int(11) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_name (name),
  KEY idx_age (age)
);

INSERT INTO student_scores
  (id, name, age, score)
VALUES
  (10, 'John', 11, 70),
  (20, 'Tom', 12, 90),
  (30, 'Jerry', 15, 95),
  (40, 'Jack', 13, 80),
  (50, 'Rose', 14, 85)
;

SET GLOBAL innodb_status_output_locks=1;
SHOW ENGINE INNODB STATUS; 
mysql> select LOCK_TYPE, LOCK_MODE, LOCK_STATUS, LOCK_DATA from performance_schema.data_locks;
```
{{< /details >}}
## Lock mode

| X/S           | X/S,REC_NOT_GAP | X/S,Gap  |
|---------------|-----------------|----------|
| Next-key lock | Row lock        | Gap lock |

## Innodb lock monitor explanation
```mysql
start transaction;
select * from student_scores where id = 20 for update;
/*
+----+------+-----+-------+
| id | name | age | score |
+----+------+-----+-------+
| 20 | Tom  |  12 |    90 |
+----+------+-----+-------+
1 row in set (0.01 sec)

RECORD LOCKS space id 159 page no 4 n bits 80 index PRIMARY of table `study`.`student_scores` trx id 12838 lock_mode X locks rec but not gap
Record lock, heap no 3 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
0: len 4; hex 80000014; asc     ;;              Primary index, value = 20 = 140x
1: len 6; hex 000000003217; asc     2 ;;
2: len 7; hex 8100000108011d; asc        ;;
3: len 3; hex 546f6d; asc Tom;;                 First column, value = Tom
4: len 1; hex 8c; asc  ;;                       Second column, value = 12 = 8c0x
5: len 4; hex 8000005a; asc    Z;;              Third column, value = 90 = 5a0x     
```

## Row lock
Lock single record. Two locks generated:
1. Intention lock granted in table level. `IX`/`IS` intention locks are `quick-test` for table level lock (X/S).
2. Record lock granted in row level.
```mysql
start transaction;
select * from student_scores where id = 20 for update;
+----+------+-----+-------+
| id | name | age | score |
+----+------+-----+-------+
| 20 | Tom  |  12 |    90 |
+----+------+-----+-------+
1 row in set (0.01 sec)

mysql> select LOCK_TYPE, LOCK_MODE, LOCK_STATUS, LOCK_DATA from performance_schema.data_locks;
+-----------+---------------+-------------+-----------+
| LOCK_TYPE | LOCK_MODE     | LOCK_STATUS | LOCK_DATA |
+-----------+---------------+-------------+-----------+
| TABLE     | IX            | GRANTED     | NULL      |
| RECORD    | X,REC_NOT_GAP | GRANTED     | 20        |
+-----------+---------------+-------------+-----------+
2 rows in set (0.01 sec)

show engine innodb status;
/*
------------
TRANSACTIONS
------------
Trx id counter 12839
Purge done for trx's n:o < 12832 undo n:o < 0 state: running but idle
History list length 0
LIST OF TRANSACTIONS FOR EACH SESSION:
---TRANSACTION 281480132365440, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 281480132364648, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 281480132363856, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 12838, ACTIVE 141 sec
2 lock struct(s), heap size 1128, 1 row lock(s)
MySQL thread id 9, OS thread handle 6122926080, query id 175 localhost root
TABLE LOCK table `study`.`student_scores` trx id 12838 lock mode IX
RECORD LOCKS space id 159 page no 4 n bits 80 index PRIMARY of table `study`.`student_scores` trx id 12838 lock_mode X locks rec but not gap
Record lock, heap no 3 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 80000014; asc     ;;              Primary index, value = 20 = 140x
 1: len 6; hex 000000003217; asc     2 ;;
 2: len 7; hex 8100000108011d; asc        ;;
 3: len 3; hex 546f6d; asc Tom;;                 First column, value = Tom
 4: len 1; hex 8c; asc  ;;                       Second column, value = 12 = 8c0x
 5: len 4; hex 8000005a; asc    Z;;              Third column, value = 90 = 5a0x
```

## Gap lock
> Especially in MySQL, there is an infimum record that is smaller than any index records of the table and a supremum record that is greater than any index records of the table.

Lock range records with one special gap locks `lock_mode X locks gap before rec`.
```mysql
mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from student_scores where id < 20 for update;
+----+-------+-----+-------+
| id | name  | age | score |
+----+-------+-----+-------+
|  5 | Nicky |   5 |    10 |
| 10 | John  |  11 |    70 |
+----+-------+-----+-------+
2 rows in set (0.01 sec)

mysql> select LOCK_TYPE, LOCK_MODE, LOCK_STATUS, LOCK_DATA from performance_schema.data_locks;
+-----------+-----------+-------------+-----------+
| LOCK_TYPE | LOCK_MODE | LOCK_STATUS | LOCK_DATA |
+-----------+-----------+-------------+-----------+
| TABLE     | IX        | GRANTED     | NULL      |
| RECORD    | X         | GRANTED     | 10        |
| RECORD    | X         | GRANTED     | 5         |
| RECORD    | X,GAP     | GRANTED     | 20        |
+-----------+-----------+-------------+-----------+
4 rows in set (0.01 sec)

mysql> show engine innodb status\G
/*
------------
TRANSACTIONS
------------
Trx id counter 12840
Purge done for trx's n:o < 12832 undo n:o < 0 state: running but idle
History list length 0
LIST OF TRANSACTIONS FOR EACH SESSION:
---TRANSACTION 281480132365440, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 281480132364648, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 281480132363856, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 12839, ACTIVE 110 sec
3 lock struct(s), heap size 1128, 3 row lock(s)
MySQL thread id 9, OS thread handle 6122926080, query id 182 localhost root
TABLE LOCK table `study`.`student_scores` trx id 12839 lock mode IX
RECORD LOCKS space id 159 page no 4 n bits 80 index PRIMARY of table `study`.`student_scores` trx id 12839 lock_mode X
Record lock, heap no 2 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 8000000a; asc     ;;
 1: len 6; hex 000000003217; asc     2 ;;
 2: len 7; hex 81000001080110; asc        ;;
 3: len 4; hex 4a6f686e; asc John;;
 4: len 1; hex 8b; asc  ;;
 5: len 4; hex 80000046; asc    F;;

Record lock, heap no 7 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 80000005; asc     ;;
 1: len 6; hex 000000003223; asc     2#;;
 2: len 7; hex 810000010c0110; asc        ;;
 3: len 5; hex 4e69636b79; asc Nicky;;
 4: len 1; hex 85; asc  ;;
 5: len 4; hex 8000000a; asc     ;;

RECORD LOCKS space id 159 page no 4 n bits 80 index PRIMARY of table `study`.`student_scores` trx id 12839 lock_mode X locks gap before rec
Record lock, heap no 3 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 80000014; asc     ;;
 1: len 6; hex 000000003217; asc     2 ;;
 2: len 7; hex 8100000108011d; asc        ;;
 3: len 3; hex 546f6d; asc Tom;;
 4: len 1; hex 8c; asc  ;;
 5: len 4; hex 8000005a; asc    Z;;
```
## Next-key lock
Differ from `gap lock` in whether owning locks for the gap itself.
```mysql
mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from student_scores where id <= 20 for update;
+----+-------+-----+-------+
| id | name  | age | score |
+----+-------+-----+-------+
|  5 | Nicky |   5 |    10 |
| 10 | John  |  11 |    70 |
| 20 | Tom   |  12 |    90 |
+----+-------+-----+-------+
3 rows in set (0.00 sec)
    
mysql> select LOCK_TYPE, LOCK_MODE, LOCK_STATUS, LOCK_DATA from performance_schema.data_locks;
+-----------+-----------+-------------+-----------+
| LOCK_TYPE | LOCK_MODE | LOCK_STATUS | LOCK_DATA |
+-----------+-----------+-------------+-----------+
| TABLE     | IX        | GRANTED     | NULL      |
| RECORD    | X         | GRANTED     | 10        |
| RECORD    | X         | GRANTED     | 20        |
| RECORD    | X         | GRANTED     | 5         |
+-----------+-----------+-------------+-----------+
4 rows in set (0.01 sec)
    
mysql> show engine innodb status\G
/*
------------
TRANSACTIONS
------------
Trx id counter 12841
Purge done for trx's n:o < 12832 undo n:o < 0 state: running but idle
History list length 0
LIST OF TRANSACTIONS FOR EACH SESSION:
---TRANSACTION 281480132365440, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 281480132364648, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 281480132363856, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 12840, ACTIVE 69 sec
2 lock struct(s), heap size 1128, 3 row lock(s)
MySQL thread id 9, OS thread handle 6122926080, query id 192 localhost root
TABLE LOCK table `study`.`student_scores` trx id 12840 lock mode IX
RECORD LOCKS space id 159 page no 4 n bits 80 index PRIMARY of table `study`.`student_scores` trx id 12840 lock_mode X
Record lock, heap no 2 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 8000000a; asc     ;;
 1: len 6; hex 000000003217; asc     2 ;;
 2: len 7; hex 81000001080110; asc        ;;
 3: len 4; hex 4a6f686e; asc John;;
 4: len 1; hex 8b; asc  ;;
 5: len 4; hex 80000046; asc    F;;

Record lock, heap no 3 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 80000014; asc     ;;
 1: len 6; hex 000000003217; asc     2 ;;
 2: len 7; hex 8100000108011d; asc        ;;
 3: len 3; hex 546f6d; asc Tom;;
 4: len 1; hex 8c; asc  ;;
 5: len 4; hex 8000005a; asc    Z;;

Record lock, heap no 7 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 80000005; asc     ;;
 1: len 6; hex 000000003223; asc     2#;;
 2: len 7; hex 810000010c0110; asc        ;;
 3: len 5; hex 4e69636b79; asc Nicky;;
 4: len 1; hex 85; asc  ;;
 5: len 4; hex 8000000a; asc     ;;
```

## Secondary index
```mysql
mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from student_scores where age > 13 for update;
+----+-------+-----+-------+
| id | name  | age | score |
+----+-------+-----+-------+
| 50 | Rose  |  14 |    85 |
| 30 | Jerry |  15 |    95 |
+----+-------+-----+-------+
2 rows in set (0.00 sec)
    
mysql> select LOCK_TYPE, LOCK_MODE, LOCK_STATUS, LOCK_DATA from performance_schema.data_locks;
+-----------+---------------+-------------+------------------------+
| LOCK_TYPE | LOCK_MODE     | LOCK_STATUS | LOCK_DATA              |
+-----------+---------------+-------------+------------------------+
| TABLE     | IX            | GRANTED     | NULL                   |
| RECORD    | X             | GRANTED     | supremum pseudo-record |
| RECORD    | X             | GRANTED     | 15, 30                 |
| RECORD    | X             | GRANTED     | 14, 50                 |
| RECORD    | X,REC_NOT_GAP | GRANTED     | 30                     |
| RECORD    | X,REC_NOT_GAP | GRANTED     | 50                     |
+-----------+---------------+-------------+------------------------+
6 rows in set (0.01 sec)

mysql> show engine innodb status\G
/*
------------
TRANSACTIONS
------------
Trx id counter 12842
Purge done for trx's n:o < 12832 undo n:o < 0 state: running but idle
History list length 0
LIST OF TRANSACTIONS FOR EACH SESSION:
---TRANSACTION 281480132365440, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 281480132364648, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 281480132363856, not started
0 lock struct(s), heap size 1128, 0 row lock(s)
---TRANSACTION 12841, ACTIVE 1276 sec
3 lock struct(s), heap size 1128, 5 row lock(s)
MySQL thread id 9, OS thread handle 6122926080, query id 198 localhost root
TABLE LOCK table `study`.`student_scores` trx id 12841 lock mode IX
RECORD LOCKS space id 159 page no 6 n bits 80 index idx_age of table `study`.`student_scores` trx id 12841 lock_mode X
Record lock, heap no 1 PHYSICAL RECORD: n_fields 1; compact format; info bits 0
 0: len 8; hex 73757072656d756d; asc supremum;;

Record lock, heap no 4 PHYSICAL RECORD: n_fields 2; compact format; info bits 0
 0: len 1; hex 8f; asc  ;;
 1: len 4; hex 8000001e; asc     ;;

Record lock, heap no 6 PHYSICAL RECORD: n_fields 2; compact format; info bits 0
 0: len 1; hex 8e; asc  ;;
 1: len 4; hex 80000032; asc    2;;

RECORD LOCKS space id 159 page no 4 n bits 80 index PRIMARY of table `study`.`student_scores` trx id 12841 lock_mode X locks rec but not gap
Record lock, heap no 4 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 8000001e; asc     ;;
 1: len 6; hex 000000003217; asc     2 ;;
 2: len 7; hex 8100000108012a; asc       *;;
 3: len 5; hex 4a65727279; asc Jerry;;
 4: len 1; hex 8f; asc  ;;
 5: len 4; hex 8000005f; asc    _;;

Record lock, heap no 6 PHYSICAL RECORD: n_fields 6; compact format; info bits 0
 0: len 4; hex 80000032; asc    2;;
 1: len 6; hex 000000003217; asc     2 ;;
 2: len 7; hex 81000001080144; asc       D;;
 3: len 4; hex 526f7365; asc Rose;;
 4: len 1; hex 8e; asc  ;;
 5: len 4; hex 80000055; asc    U;;
```
## Reference
[understand-the-basics-of-locks-and-deadlocks-in-mysql-part-i](https://lynn-kwong.medium.com/understand-the-basics-of-locks-and-deadlocks-in-mysql-part-i-92f229db0a)
[mysql-deadlocks](https://github.com/aneasystone/mysql-deadlocks)