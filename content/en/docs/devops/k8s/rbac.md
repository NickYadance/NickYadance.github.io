---
title: "Rbac"
description: ""
lead: ""
date: 2022-09-15T22:00:13+08:00
lastmod: 2022-09-15T22:00:13+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "rbac-39a2a98e2291f19e2f143e704d02db9a"
weight: 999
toc: true
---
Serviceaccount with admin role for quick account setup.
```yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: airflow-worker
  namespace: airflow
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: airflow-worker
  namespace: airflow # Should be namespace you are granting access to
rules:
  - apiGroups: [""]
    resources: ["pods", "pods/log"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["pods/exec"]
    verbs: ["create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: airflow-worker
  namespace: airflow # Should be namespace you are granting access to
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: airflow-worker # Should match name of Role
subjects:
  - namespace: airflow
    kind: ServiceAccount
    name: airflow-worker # Should match service account name, above
```