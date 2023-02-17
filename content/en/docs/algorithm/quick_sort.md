---
title: "Quicksort"
description: ""
lead: ""
date: 2023-02-11T21:47:50+08:00
lastmod: 2023-02-11T21:47:50+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "quick_sort-16c9b8ebfc9edff1de4e210b53f7c897"
weight: 999
toc: true
---
## Divide and Conquer
```text
QuickSort(A, l, r):
     *  if l < r :
     *      p = Partition(A, l, r)
     *      QuickSort(A, l, p - 1)
     *      QuickSort(A, p + 1, r)
```

## Hoare Partition and Lomuto Partition
[hoares-vs-lomuto-partition-scheme-quicksort](https://www.geeksforgeeks.org/hoares-vs-lomuto-partition-scheme-quicksort/)

In Lomuto partition, the array is divided into 4 parts.
![img.png](images/quicksort_partition.png)

For $A[l, r), i \in [l, r), j \in [l, r), x=A[r - 1]$
$$
\begin{cases}
V: &T(j)= \begin{cases} &k \in [l, i), A[k] <= x \\\\ &k \in (i, j), A[k] > x \\\\ &k \in [j, r), A[k] \ ? \ x \end{cases} \\\\
S: &T(j)= i = l, j = l \\\\
L: &T(j+1)= A[j] <= x \  ? \ swap(A, i++, j) \\\\
E: &Answer= \begin{cases} &k \in [l, i), A[k] <= x \\\\ &k \in (i, r), A[k] > x \\\\ &A[i] = x \end{cases} \\\\
\end{cases}
$$

The $Loop$ swaps A[j] to A[i] if A[j] <= x so the variant keeps.

![img.png](images/quicksort_loop.png)
## Pivot
The $x$ in the loop variants is called the `pivot`. There are several ways to choose it and the
point is to have a more **randomized** pick to produce balanced sub results.

## Performance
The reason being called quicksort is that the quicksort has average performance at $O(nlgn)$.

The partition, obviously cost O(n) for each travel.
The cost of divide and conquer, however with related to how balanced the array is divided.

The detailed provence will be complex, conclusion is that in most conditions, the partition
will produce rather balanced results. That is why it is quicker than other sort algorithms.

## Code
{{< details "Quicksort" >}}
```java
  public static void quickSort(int[] A) {
        quickSort(A, 0, A.length);
    }

    public static void quickSort(int[] A, int l, int r) {
        if (l < r) {
            int p = partition(A, l, r);
            quickSort(A, l, p);
            quickSort(A, p + 1, r);
        }
    }

    public static int partition(int[] A, int l, int r) {
        int x = A[r - 1];
        int i = l, j = l;
        while (j < r) {
            if (A[j] <= x) {
                swap(A, i++, j);
            }
            ++j;
        }
        return i - 1;
    }
```
{{< /details >}}
