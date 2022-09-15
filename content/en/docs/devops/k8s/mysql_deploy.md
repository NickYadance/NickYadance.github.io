---
title: "Mysql Deploy"
description: ""
lead: ""
date: 2022-09-12T18:31:51+08:00
lastmod: 2022-09-12T18:31:51+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "mysql_deploy-3ab4dd33ad50194f229e953d6c0b31d3"
weight: 999
toc: true
---
Single node mysql deployment.
```yaml
---
apiVersion: v1                    # API version
kind: Service                     # Type of kubernetes resource
metadata:
  name: mysql         # Name of the resource
  labels:                         # Labels that will be applied to the resource
    app: mysql
  namespace: airflow
spec:
  ports:
    - port: 3306
  selector:
    app: mysql
  clusterIP: 10.96.144.158
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql-config
  namespace: airflow
# Extra configurations
#    set innodb_lock_wait_timeout=100;
data:
  my.cnf: |-
    [mysqld]
    max_connections  = 4096
    binlog_expire_logs_seconds = 3600
    explicit_defaults_for_timestamp = 1
---

apiVersion: apps/v1
kind: Deployment                    # Type of the kubernetes resource
metadata:
  name: mysql           # Name of the deployment
  labels:                           # Labels applied to this deployment
    app: mysql
  namespace: airflow
spec:
  selector:
    matchLabels:                    # This deployment applies to the Pods matching the specified labels
      app: mysql
  strategy:
    type: Recreate
  template:                         # Template for the Pods in this deployment
    metadata:
      labels:                       # Labels to be applied to the Pods in this deployment
        app: mysql
    spec:                           # The spec for the containers that will be run inside the Pods in this deployment
      nodeSelector:
        database: accept
      tolerations:
        - key: node-role.kubernetes.io/master
          operator: Equal
          effect: NoSchedule
      containers:
        - image: mysql:8            # The container image
          name: mysql
          securityContext:
            runAsUser: 0
          env:                        # Environment variables passed to the container
            - name: MYSQL_ROOT_PASSWORD
              valueFrom:                # Read environment variables from kubernetes secrets
                secretKeyRef:
                  name: mysql-root-pass-96h4525584
                  key: password
            - name: MYSQL_DATABASE
              valueFrom:
                secretKeyRef:
                  name: mysql-db-url-cdthkb6gg6
                  key: database
            - name: MYSQL_USER
              valueFrom:
                secretKeyRef:
                  name: mysql-user-pass-2bg297f5c9
                  key: username
            - name: MYSQL_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: mysql-user-pass-2bg297f5c9
                  key: password
          ports:
            - containerPort: 3306        # The port that the container exposes
              name: mysql
          volumeMounts:
            - name: mysql-persistent-storage  # This name should match the name specified in `volumes.name`
              mountPath: /var/lib/mysql
            - name: mysql-config
              mountPath: /etc/mysql/conf.d
      volumes:
        - name: mysql-persistent-storage
          hostPath:
            path: /mnt/airflow/data
        - name: mysql-config
          configMap:
            name: mysql-config
```