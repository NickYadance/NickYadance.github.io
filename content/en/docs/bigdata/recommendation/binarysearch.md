---
title: "Binarysearch"
description: ""
lead: ""
date: 2023-01-11T10:13:09+08:00
lastmod: 2023-01-11T10:13:09+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "binarysearch-bd5bfbf2190dc966dae56ab09c002e6c"
weight: 999
toc: true
---
## lower_bound
lower_bound of $x$ in non-descending $array[l, r)$ is the leftmost index to insert $x$ while keeping order.

To find the lower_bound with binary-search, keep loop invariant as below during iterations.
$$
loop\ invariant
\begin{cases}
0.\\ [l, r), &l<r \\\\
1.\\ a[i]<=a[j], &l<=i<j<r \\\\
2.\\ array[l_0]<array[l], &l_0<l \\\\
3.^*\\ array[r_0]>=array[r], &r_0>=r \\\\
\end{cases}
$$
The invariant $3.$ can be inferred from $0,1,2$.

Basic implementation in Java.
```java
class BinarySearch {
  public static int lower_bound(int[] a, int target) {
    int l = 0, r = a.length;
    while (l < r) {
      int mid = (l + r) / 2;
      if (a[mid] < target) {
        l = mid + 1;
      } else {
        r = mid;
      }
    }
    return l;
  }
}
```
## upper_bound
upper_bound of $x$ in non-descending $array[l, r)$ is the rightmost index to insert $x$ while keeping order.
$$
loop\ invariant
\begin{cases}
0.\\ [l, r), &l<r \\\\
1.\\ a[i]<=a[j], &l<=i<j<r \\\\
2.\\ array[l_0]<=array[l], &l_0<l \\\\
3.^*\\ array[r_0]>array[r], &r_0>=r \\\\
\end{cases}
$$
The invariant $3.$ can be inferred from $0,1,2$. Specially for int array $$upper\\_bound(x)=lower\\_bound(x+1) - 1$$

Note that when we return $l$ in upper_bound, it is the index to **insert before**, so for $array[1,2,3]$ and
$x=2$, we return 2 instead of 1 cauze we need to insert x before.
```java
class BinarySearch {
  public static int upper_bound(int[] a, int target) {
    int l = 0, r = a.length;
    while (l < r) {
      int mid = (l + r) / 2;
      if (a[mid] <= target) {
        l = mid + 1;
      } else {
        r = mid;
      }
    }
    return l;
  }
}
```
