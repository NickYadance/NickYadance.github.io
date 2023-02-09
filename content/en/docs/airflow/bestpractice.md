---
title: "Bestpractice"
description: ""
lead: ""
date: 2022-11-22T11:06:19+08:00
lastmod: 2022-11-22T11:06:19+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "bestpractice-3097e437df321bf19b2e5754f3628f7f"
weight: 999
toc: true
---
## Take care Database, especially Mysql
* Connections
* IO : Airflow has heavily read ops and mid-level write ops 
* Slow query : Airflow can make LOTS of slow queries even if the database is properly configured
* Clean history data 

## Disable statsd if you can
* statsd query is large, which can bring extra pressure to the db
* statsd can potentially crush your scheduler ([#18010](https://github.com/apache/airflow/issues/18010))

## Patch critical section of scheduler
* Patch this ([PR](https://github.com/apache/airflow/pull/25673)) if Airflow version <2.3.4. This can really slow Airflow when dataset is large
* Tune scheduler configuration constantly 

## Let Airflow do the scheduling
* Use Celery or K8s to save the heavy stuff from Airflow scheduler
