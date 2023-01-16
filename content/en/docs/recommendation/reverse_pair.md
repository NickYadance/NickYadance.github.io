---
title: "Reverse Pair"
description: ""
lead: ""
date: 2023-01-11T17:34:48+08:00
lastmod: 2023-01-11T17:34:48+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "reverse_pair-7dc625a8d5919a087865c9e82bf20d1e"
weight: 999
toc: true
---
> What's the relationship between O(n) of insertion sort and reverse pair ?

A: Equal.

```java
class ReversePair {
  public static int reversePair(int[] A, int l, int r) {
    if (r - l > 1) {
      int mid = (r + l) / 2;
      int n = 0;
      for (int i = l; i < mid; i++) {
        for (int j = mid; j < r; j++) {
          if (A[i] > A[j]) ++n;
        }
      }
      return n + reversePair(A, l, mid) + reversePair(A, mid, r);
    }
    return 0;
  }
}
```
