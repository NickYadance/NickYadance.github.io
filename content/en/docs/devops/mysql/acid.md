---
title: "ACID"
description: ""
lead: ""
date: 2022-09-27T16:15:41+08:00
lastmod: 2022-09-27T16:15:41+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "acid-3fe482921224f31913a33f84e11f943e"
weight: 999
toc: true
---
## Data
```mysql
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
UNION ALL SELECT 1,2,3
;
```
## Read Uncommitted
```mysql
-- T1                                              - T2
SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;  SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
BEGIN;                                             BEGIN;
UPDATE IsolationTests SET Col1 = 2;                -- Dirty Read. 
DO sleep(10);                                      SELECT * FROM IsolationTests; 
ROLLBACK;                                          COMMIT; 
```
## Read Committed
> If the transaction isolation level is REPEATABLE READ (the default level), all consistent reads within the same transaction read the snapshot established **by the first such read in that transaction.** You can get a fresher snapshot for your queries by committing the current transaction and after that issuing new queries. 
> With READ COMMITTED isolation level, each consistent read within a transaction **sets and reads its own fresh snapshot.**

> This is called multi-versioned concurrency control. It increases transaction concurrency by split read from write and reduce locked read.

```mysql
-- T1                                             -- T2                                            
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;   SET TRANSACTION ISOLATION LEVEL READ COMMITTED;  
BEGIN;                                            BEGIN;                                           
                                                  -- Before T1.update : Col1 = 1                   
                                                  SELECT * FROM IsolationTests;                    
UPDATE IsolationTests SET Col1 = 2;               
                                                  -- After T1.update : Col1 = 2. Unrepeatable read 
DO sleep(10);                                     SELECT * FROM IsolationTests;                    
ROLLBACK;                                         COMMIT;                                          
```                                               
## Repeatable Read (Default)
```mysql
-- T1                                              -- T2                                            
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ ;  SET TRANSACTION ISOLATION LEVEL REPEATABLE READ ;
BEGIN;                                             BEGIN;                                           
                                                   -- Before T1.update : Col1 = 1
                                                   SELECT * FROM IsolationTests;  
UPDATE IsolationTests SET Col1 = 2;                                   
DO sleep(10);                                      -- After T1.update : Col1 = 1                     
ROLLBACK;                                          SELECT * FROM IsolationTests;                    
                                                   COMMIT; 
```            
## Repeatable Read (Phantom Read)
```mysql
-- T1                                              -- T2                                            
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ ;  SET TRANSACTION ISOLATION LEVEL REPEATABLE READ ;
BEGIN;                                             BEGIN;                                           
                                                   -- Before T1.commit : empty set
                                                   SELECT * FROM IsolationTests where Col1=2;  
INSERT INTO IsolationTests values (2,2,3)                                 
COMMIT;                                            -- After T1.commit : empty set     
                                                   SELECT * FROM IsolationTests where Col1=2;   
                                                   -- Update column, forced latest snapshot : success
                                                   UPDATE IsolationTests SET Col2=3, Col3=4 where Col1=2;
                                                   -- Phantom Read : Col (2, 3, 4)
                                                   SELECT * FROM IsolationTests where Col1=2;
                                                   COMMIT; 
```
## Reference
[innodb-consistent-read](https://dev.mysql.com/doc/refman/5.7/en/innodb-consistent-read.html)

[sql-server-isolation-levels-by-example](https://gavindraper.com/2012/02/18/sql-server-isolation-levels-by-example/)