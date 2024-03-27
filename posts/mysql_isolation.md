---
title: 'Mysql事务隔离级别(无废话版)'
date: '2023-12-18'
description: 'Mysql事务隔离级别(无废话版)'
---

测试数据
```nsql
DROP TABLE IF EXISTS IsolationTests;
CREATE TABLE IsolationTests
(
    Col1 INT,
    Col2 INT,
    Col3 INT
);

INSERT INTO IsolationTests(Col1,Col2,Col3)
SELECT 1,2,3
UNION ALL SELECT 1,2,3
UNION ALL SELECT 1,2,3
UNION ALL SELECT 1,2,3
UNION ALL SELECT 1,2,3
UNION ALL SELECT 1,2,3
UNION ALL SELECT 1,2,3;
```

## 读未提交(READ UNCOMMITTED)
```nsql
-- T1                                              - T2
SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;  SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
BEGIN;                                             BEGIN;
UPDATE IsolationTests SET Col1 = 2;                
                                                   -- Col1=2, dirty read 
                                                   SELECT * FROM IsolationTests;    
ROLLBACK;                                          
                                                   -- Col1=1
                                                   SELECT * FROM IsolationTests;  
                                                   COMMIT; 
```

## 读已提交(READ COMMITTED)
```nsql
-- T1                                             -- T2                                            
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;   SET TRANSACTION ISOLATION LEVEL READ COMMITTED;  
BEGIN;                                            BEGIN;                                           
                                                  -- Col1 = 1                   
                                                  SELECT * FROM IsolationTests;    
UPDATE IsolationTests SET Col1 = 2;               
                                                  -- Col1 = 1, T1 uncommitted 
                                                  SELECT * FROM IsolationTests;    
COMMIT;
                                                  -- Col1 = 2, T1 committed 
                                                  SELECT * FROM IsolationTests;    
												  COMMIT;                                          
```

## 可重复读(REPEATABLE READ)
```nsql
-- T1                                             -- T2                                            
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;  SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;  
BEGIN;                                            BEGIN;                                           
                                                  -- Col1 = 1                   
                                                  SELECT * FROM IsolationTests;    
UPDATE IsolationTests SET Col1 = 2;               
                                                  -- Col1 = 1, T1 uncommitted
                                                  SELECT * FROM IsolationTests;    
COMMIT;
                                                  -- Col1 = 1, T1 committed
                                                  SELECT * FROM IsolationTests;    
                                                  COMMIT;
```

## 幻读(REPEATABLE READ)
```nsql
-- T1                                              -- T2                                            
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ ;  SET TRANSACTION ISOLATION LEVEL REPEATABLE READ ;
BEGIN;                                             BEGIN;                                           
                                                   -- empty 
                                                   SELECT * FROM IsolationTests where Col1=2;  
INSERT INTO IsolationTests values (2,2,3);                                 
COMMIT;                                            
                                                   -- empty      
                                                   SELECT * FROM IsolationTests where Col1=2;   
                                                   -- Update with forced latest snapshot 
                                                   UPDATE IsolationTests SET Col2=3, Col3=4 where Col1=2;
                                                   -- Phantom Read : Col (2, 3, 4)
                                                   SELECT * FROM IsolationTests where Col1=2;
                                                   COMMIT; 
```


> * [innodb-consistent-read](https://dev.mysql.com/doc/refman/5.7/en/innodb-consistent-read.html)  
> * [sql-server-isolation-levels-by-example](https://gavindraper.com/2012/02/18/sql-server-isolation-levels-by-example/)  